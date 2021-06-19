package service

import "gitabza-go/model"

type AddNewArtistRequest struct {
	Name    string `json:"name" binding:"required"`
	Slug    string `json:"slug" binding:"required"`
	Picture string `json:"picture"`
}

func (req *AddNewArtistRequest) ToModel() *model.Artist {
	return &model.Artist{
		Name:    &req.Name,
		Slug:    &req.Slug,
		Picture: &req.Picture,
	}
}
