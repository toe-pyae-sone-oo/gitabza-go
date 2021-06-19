package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Artist struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Name      *string            `bson:"name, omitempty"`
	Slug      *string            `bson:"slug, omitempty"`
	Picture   *string            `bson:"picture, omitempty"`
	UUID      *string            `bson:"uuid, omitempty"`
	CreatedAt *time.Time         `bson:"created_at, omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at, omitempty"`
}

func (a *Artist) ToBson() primitive.M {
	return bson.M{
		"name":       a.Name,
		"slug":       a.Slug,
		"picture":    a.Picture,
		"uuid":       a.UUID,
		"created_at": a.CreatedAt,
		"updated_at": a.UpdatedAt,
	}
}
