package service

import "gitabza-go/model"

type AddNewArtistResponse struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Picture string `json:"picture"`
}

func (resp *AddNewArtistResponse) FromModel(md *model.Artist) {
	resp.UUID = *md.UUID
	resp.Name = *md.Name
	resp.Slug = *md.Slug
	resp.Picture = *md.Picture
}
