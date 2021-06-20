package repository

import (
	"context"
	"gitabza-go/model"
	"gitabza-go/mongodb"

	"go.mongodb.org/mongo-driver/bson"
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
