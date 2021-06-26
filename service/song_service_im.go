package service

import (
	"context"
	"gitabza-go/repository"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type SongServiceIM struct {
	songRepo repository.SongRepository
}

func NewSongService() SongService {
	return &SongServiceIM{
		songRepo: repository.NewSongMongoRepository(),
	}
}

func (s *SongServiceIM) FindByArtist(ctx context.Context, artistID string) (SongsResponse, error) {
	songs, err := s.songRepo.FindByArtist(ctx, artistID)
	if err != nil {
		return nil, err
	}

	var resp SongsResponse
	resp.FromModel(songs)

	return resp, nil
}

func (s *SongServiceIM) Add(ctx context.Context, req *AddNewSongRequest) (*AddNewSongResponse, error) {
	if req == nil {
		return nil, errors.Wrap(ErrBadRequest, "req must not be nil")
	}

	foundSong, err := s.songRepo.FindOneBySlug(ctx, req.Slug)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if foundSong != nil {
		return nil, errors.Wrapf(ErrBadRequest, "song with slug %s already exists", req.Slug)
	}

	newSong := req.ToModel()
	if err := s.songRepo.Add(ctx, newSong); err != nil {
		return nil, err
	}

	resp := new(AddNewSongResponse)
	resp.FromModel(newSong)

	return resp, nil
}

func (s *SongServiceIM) Find(ctx context.Context, q *FindSongsQuery) (SongsResponse, error) {
	songs, err := s.songRepo.SearchByTitle(ctx, q.Title,
		int64(q.Skip), int64(q.Limit), q.Sort, q.Order)
	if err != nil {
		return nil, err
	}

	var resp SongsResponse
	resp.FromModel(songs)
	return resp, nil
}

func (s *SongServiceIM) Delete(ctx context.Context, uuid string) error {
	return s.songRepo.DeleteOneByUUID(ctx, uuid)
}
