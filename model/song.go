package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	Artists    []string           `bson:"artists"`
	Title      *string            `bson:"title, omitempty"`
	Slug       *string            `bson:"slug, omitempty"`
	Genre      *string            `bson:"genre, omitempty"`
	Stype      *string            `bson:"types, omitempty"`
	Difficulty *string            `bson:"difficulty, omitempty"`
	Capo       *string            `bson:"capo, omitempty"`
	Version    *string            `bson:"version, omitempty"`
	Lyrics     *string            `bson:"lyrics, omitempty"`
	Youtube    *string            `bson:"youtube, omitempty"`
	UUID       *string            `bson:"uuid, omitempty"`
	CreatedAt  *time.Time         `bson:"created_at, omitempty"`
	UpdatedAt  *time.Time         `bson:"updated_at, omitempty"`
}
