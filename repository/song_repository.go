package repository

import (
	"context"
	"gitabza-go/model"
)

type SongRepository interface {
	FindByArtist(ctx context.Context, artistID string) ([]model.Song, error)
}
