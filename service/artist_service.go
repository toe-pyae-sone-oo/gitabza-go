package service

import (
	"context"
)

type ArtistService interface {
	Add(ctx context.Context, req *AddNewArtistRequest) (*AddNewArtistResponse, error)
}
