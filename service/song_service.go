package service

import "context"

type SongService interface {
	FindByArtist(ctx context.Context, artistID string) (SongsResponse, error)
	Add(ctx context.Context, req *AddNewSongRequest) (*AddNewSongResponse, error)
	Find(ctx context.Context, q *FindSongsQuery) (SongsResponse, error)
}
