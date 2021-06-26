package repository

import (
	"context"
	"gitabza-go/model"
)

type SongRepository interface {
	FindByArtist(ctx context.Context, artistID string) ([]model.Song, error)
	FindOneBySlug(ctx context.Context, slug string) (*model.Song, error)
	Add(ctx context.Context, song *model.Song) error
	SearchByTitle(ctx context.Context, title string, skip, limit int64, sortBy, orderBy string) ([]model.Song, error)
}
