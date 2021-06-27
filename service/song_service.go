package service

import "context"

type SongService interface {
	FindByArtist(ctx context.Context, artistID string) (SongsResponse, error)
	Add(ctx context.Context, req *AddNewSongRequest) (*AddNewSongResponse, error)
	Find(ctx context.Context, q *FindSongsQuery) (SongsResponse, error)
	Delete(ctx context.Context, uuid string) error
	FindBySlug(ctx context.Context, slug string) (*SongResponse, error)
	IsSongRelatedToArtist(ctx context.Context, song *SongResponse, artistSlug string) (bool, error)
}
