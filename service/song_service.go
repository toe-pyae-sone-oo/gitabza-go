package service

import "context"

type SongService interface {
	FindByArtist(ctx context.Context, artistID string) (SongsResponse, error)
}
