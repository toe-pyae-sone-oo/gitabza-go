package service

import (
	"context"
)

type ArtistService interface {
	Add(ctx context.Context, req *AddNewArtistRequest) (*AddNewArtistResponse, error)
	Find(ctx context.Context, query *FindArtistsQuery) (ArtistsResponse, error)
	Delete(ctx context.Context, uuid string) error
}
