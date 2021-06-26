package repository

import (
	"context"
	"gitabza-go/model"
)

type SongRepository interface {
	FindByArtist(ctx context.Context, artistID string) ([]model.Song, error)
	FindOneBySlug(ctx context.Context, slug string) (*model.Song, error)
	Add(ctx context.Context, song *model.Song) error
}
