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

type SongMongoRepository struct {
	db            *mongo.Client
	coll          *mongo.Collection
	excludeFields interface{}
}

func NewSongMongoRepository() SongRepository {
	_db := mongodb.GetDB()
	return &SongMongoRepository{
		db:   _db,
		coll: _db.Database(DbName).Collection(ColSongs),
		excludeFields: bson.D{
			{Key: "_id", Value: 0},
			{Key: "updated_at", Value: 0},
		},
	}
}

func (r *SongMongoRepository) FindByArtist(ctx context.Context, artistID string) ([]model.Song, error) {
	filter := bson.M{"artists": artistID}
	opts := options.Find().
		SetSort(bson.M{"title": 1}).
		SetProjection(r.excludeFields)

	var results []model.Song
	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *SongMongoRepository) FindOneBySlug(ctx context.Context, slug string) (*model.Song, error) {
	filter := bson.M{"slug": slug}
	opts := options.FindOne().SetProjection(r.excludeFields)

	song := new(model.Song)
	if err := r.coll.FindOne(ctx, filter, opts).Decode(song); err != nil {
		return nil, err
	}

	return song, nil
}

func (r *SongMongoRepository) Add(ctx context.Context, song *model.Song) error {
	if song == nil {
		return errors.Wrap(ErrInvalidParm, "song must not be nil")
	}

	song.UUID = typeutil.String(uuid.NewString())
	song.CreatedAt = typeutil.Time(time.Now())
	song.UpdatedAt = typeutil.Time(time.Now())

	res, err := r.coll.InsertOne(ctx, song.ToBson())
	if err != nil {
		return err
	}

	song.ID = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (r *SongMongoRepository) SearchByTitle(
	ctx context.Context, title string,
	skip, limit int64,
	sortBy, orderBy string,
) ([]model.Song, error) {
	filter := bson.M{"title": primitive.Regex{Pattern: title, Options: "i"}}

	desc := -1
	if orderBy == "asc" {
		desc = 1
	}

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.M{sortBy: desc}).
		SetProjection(r.excludeFields)

	var results []model.Song
	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *SongMongoRepository) DeleteOneByUUID(ctx context.Context, uuid string) error {
	filter := bson.M{"uuid": uuid}
	if _, err := r.coll.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
