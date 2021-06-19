package repository

import (
	"context"
	"gitabza-go/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArtistRepository interface {
	FindOneBySlug(ctx context.Context, slug string) (*model.Artist, error)
	SearchByName(ctx context.Context, name string, skip, limit int, sortBy, orderBy string) ([]model.Artist, error)
	FindOneByUUID(ctx context.Context, uuid string) (*model.Artist, error)
	DeleteOneByUUID(ctx context.Context, uuid string) error
	GetAllNames(ctx context.Context) ([]string, error)
	Add(ctx context.Context, artist *model.Artist) error
	UpdateByUUID(ctx context.Context, uuid string, payload *model.Artist) (*model.Artist, error)
	FindOneByID(ctx context.Context, id primitive.ObjectID) (*model.Artist, error)
}
