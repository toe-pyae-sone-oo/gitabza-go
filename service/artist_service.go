package service

import (
	"context"
)

type ArtistService interface {
	Add(ctx context.Context, req *AddNewArtistRequest) (*AddNewArtistResponse, error)
	Find(ctx context.Context, query *FindArtistsQuery) (ArtistsResponse, error)
	Delete(ctx context.Context, uuid string) error
	GetAllNames(ctx context.Context) ([]string, error)
	FindBySlug(ctx context.Context, slug string) (*ArtistResponse, error)
	FindByUUID(ctx context.Context, uuid string) (*ArtistResponse, error)
	Update(ctx context.Context, req *UpdateArtistRequest) (*UpdateArtistResponse, error)
}
