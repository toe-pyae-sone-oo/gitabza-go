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
