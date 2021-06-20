package service

import (
	"context"
	"gitabza-go/repository"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArtistServiceIM struct {
	artistRepo repository.ArtistRepository
}

func NewArtistService() ArtistService {
	return &ArtistServiceIM{
		artistRepo: repository.NewArtistMongoRepository(),
	}
}

func (s *ArtistServiceIM) Add(ctx context.Context, req *AddNewArtistRequest) (*AddNewArtistResponse, error) {
	if req == nil {
		return nil, errors.Wrap(ErrBadRequest, "req must not be null")
	}

	foundArtist, err := s.artistRepo.FindOneBySlug(ctx, req.Slug)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if foundArtist != nil {
		return nil, errors.Wrapf(ErrBadRequest, "artist with slug %s already exists", req.Slug)
	}

	newArtist := req.ToModel()
	if err := s.artistRepo.Add(ctx, newArtist); err != nil {
		return nil, err
	}

	resp := new(AddNewArtistResponse)
	resp.FromModel(newArtist)

	return resp, nil
}

func (s *ArtistServiceIM) Find(ctx context.Context, query *FindArtistsQuery) (ArtistsResponse, error) {
	artists, err := s.artistRepo.SearchByName(ctx, query.Name,
		int(query.Skip), int(query.Limit), query.Sort, query.Order)
	if err != nil {
		return nil, err
	}

	var resp ArtistsResponse
	resp.FromModel(artists)
	return resp, nil
}

func (s *ArtistServiceIM) Delete(ctx context.Context, uuid string) error {
	return s.artistRepo.DeleteOneByUUID(ctx, uuid)
}

func (s *ArtistServiceIM) GetAllNames(ctx context.Context) ([]string, error) {
	return s.artistRepo.GetAllNames(ctx)
}

func (s *ArtistServiceIM) FindBySlug(ctx context.Context, slug string) (*ArtistResponse, error) {
	artist, err := s.artistRepo.FindOneBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	var resp ArtistResponse
	resp.FromModel(artist)
	return &resp, nil
}
