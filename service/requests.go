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

type FindArtistsQuery struct {
	Name  string `form:"name"`
	Skip  uint   `form:"skip"`
	Limit uint   `form:"limit,default=10"`
	Sort  string `form:"sort,default='created_at'"`
	Order string `form:"order,default='desc'"`
}
