package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Song) ToBson() primitive.M {
	return bson.M{
		"artists":    s.Artists,
		"title":      s.Title,
		"slug":       s.Slug,
		"types":      s.Stype,
		"difficulty": s.Difficulty,
		"capo":       s.Capo,
		"version":    s.Version,
		"lyrics":     s.Lyrics,
		"youtube":    s.Youtube,
		"uuid":       s.UUID,
		"created_at": s.CreatedAt,
		"updated_at": s.UpdatedAt,
	}
}
