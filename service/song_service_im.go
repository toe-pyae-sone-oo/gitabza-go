package service

import (
	"context"
	"gitabza-go/repository"
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
