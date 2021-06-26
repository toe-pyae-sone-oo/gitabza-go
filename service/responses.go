package service

import "gitabza-go/model"

type AddNewArtistResponse struct {
	ArtistResponse
}

type ArtistResponse struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Picture string `json:"picture"`
}

func (resp *ArtistResponse) FromModel(md *model.Artist) {
	resp.UUID = *md.UUID
	resp.Name = *md.Name
	resp.Slug = *md.Slug
	resp.Picture = *md.Picture
}

type ArtistsResponse []ArtistResponse

func (resp *ArtistsResponse) FromModel(mds []model.Artist) {
	*resp = make([]ArtistResponse, len(mds))
	for i, md := range mds {
		(*resp)[i].FromModel(&md)
	}
}

type UpdateArtistResponse struct {
	ArtistResponse
}

type SongResponse struct {
	UUID       string `json:"uuid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Stype      string `json:"types"`
	Difficulty string `json:"difficulty"`
	Capo       string `json:"capo"`
	Version    string `json:"version"`
	Lyrics     string `json:"lyrics"`
	Youtube    string `json:"youtube"`
}

func (resp *SongResponse) FromModel(md *model.Song) {
	resp.UUID = *md.UUID
	resp.Title = *md.Title
	resp.Slug = *md.Slug
	resp.Stype = *md.Stype
	resp.Difficulty = *md.Difficulty
	resp.Capo = *md.Capo
	resp.Version = *md.Version
	resp.Lyrics = *md.Lyrics
	resp.Youtube = *md.Youtube
}

type SongsResponse []SongResponse

func (resp *SongsResponse) FromModel(mds []model.Song) {
	*resp = make([]SongResponse, len(mds))
	for i, md := range mds {
		(*resp)[i].FromModel(&md)
	}
}

type AddNewSongResponse struct {
	SongResponse
}
