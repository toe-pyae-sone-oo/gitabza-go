package repository

import (
	"context"
	"gitabza-go/common/typeutil"
	"gitabza-go/model"
	"gitabza-go/mongodb"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArtistMongoRepository struct {
	db            *mongo.Client
	coll          *mongo.Collection
	excludeFields interface{}
}

func NewArtistMongoRepository() ArtistRepository {
	_db := mongodb.GetDB()
	return &ArtistMongoRepository{
		db:   _db,
		coll: _db.Database("gitabza").Collection("artists"),
		excludeFields: bson.D{
			{Key: "_id", Value: 0},
			{Key: "updated_at", Value: 0},
		},
	}
}

func (r *ArtistMongoRepository) FindOneBySlug(ctx context.Context, slug string) (*model.Artist, error) {
	filter := bson.M{"slug": slug}
	opts := options.FindOne().SetProjection(r.excludeFields)

	artist := &model.Artist{}
	if err := r.coll.FindOne(ctx, filter, opts).Decode(artist); err != nil {
		return nil, err
	}

	return artist, nil
}

// this method is costly. should later modify using search index
// or migrate to elasticsearch
func (r *ArtistMongoRepository) SearchByName(
	ctx context.Context, name string,
	skip, limit int,
	sortBy, orderBy string,
) ([]model.Artist, error) {
	filter := bson.M{"name": primitive.Regex{Pattern: name, Options: "i"}}

	desc := -1
	if orderBy == "asc" {
		desc = 1
	}

	opts := options.Find().
		SetSort(bson.M{"created_at": 1}).
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.M{sortBy: desc}).
		SetProjection(r.excludeFields)

	var results []model.Artist

	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return results, err
	}

	if err := cur.All(ctx, &results); err != nil {
		return results, err
	}

	return results, nil
}

func (r *ArtistMongoRepository) FindOneByUUID(ctx context.Context, uuid string) (*model.Artist, error) {
	filter := bson.M{"uuid": uuid}
	opts := options.FindOne().SetProjection(r.excludeFields)

	artist := &model.Artist{}
	if err := r.coll.FindOne(ctx, filter, opts).Decode(artist); err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *ArtistMongoRepository) DeleteOneByUUID(ctx context.Context, uuid string) error {
	filter := bson.M{"uuid": uuid}

	if _, err := r.coll.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (r *ArtistMongoRepository) GetAllNames(ctx context.Context) ([]string, error) {
	opts := options.Find().SetProjection(bson.M{"name": 1})

	names := make([]string, 0)
	artists := make([]model.Artist, 0)

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return names, err
	}

	if err := cur.All(ctx, &artists); err != nil {
		return names, err
	}

	for _, a := range artists {
		names = append(names, *a.Name)
	}

	return names, nil
}

func (r *ArtistMongoRepository) Add(ctx context.Context, artist *model.Artist) error {
	if artist == nil {
		return errors.Wrap(ErrInvalidParm, "artist must not be null")
	}

	artist.UUID = typeutil.String(uuid.NewString())
	artist.CreatedAt = typeutil.Time(time.Now())
	artist.UpdatedAt = typeutil.Time(time.Now())

	res, err := r.coll.InsertOne(ctx, artist.ToBson())
	if err != nil {
		return err
	}

	artist.ID = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (r *ArtistMongoRepository) UpdateByUUID(ctx context.Context, uuid string, payload *model.Artist) (*model.Artist, error) {
	if payload == nil {
		return nil, errors.Wrap(ErrInvalidParm, "payload must not be null")
	}
	payload.UpdatedAt = typeutil.Time(time.Now())

	pb := payload.ToBson()
	delete(pb, "created_at")
	delete(pb, "uuid")

	if err := r.coll.FindOneAndUpdate(ctx, bson.M{"uuid": uuid}, bson.M{"$set": pb}).Err(); err != nil {
		return nil, err
	}

	return r.FindOneByUUID(ctx, uuid)
}

func (r *ArtistMongoRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*model.Artist, error) {
	filter := bson.M{"_id": id}
	opts := options.FindOne().SetProjection(r.excludeFields)

	artist := &model.Artist{}
	if err := r.coll.FindOne(ctx, filter, opts).Decode(artist); err != nil {
		return nil, err
	}

	return artist, nil
}
