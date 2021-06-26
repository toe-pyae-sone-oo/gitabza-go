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

type UpdateArtistRequest struct {
	UUID    string
	Name    string  `json:"name" binding:"required"`
	Slug    string  `json:"slug" binding:"required"`
	Picture *string `json:"picture"`
}

func (req *UpdateArtistRequest) ToModel() *model.Artist {
	return &model.Artist{
		Name:    &req.Name,
		Slug:    &req.Slug,
		Picture: req.Picture,
	}
}

type AddNewSongRequest struct {
	Title      string   `json:"title" binding:"required"`
	Slug       string   `json:"slug" binding:"required"`
	Artists    []string `json:"artists" binding:"required"` // check this
	Stype      string   `json:"types" binding:"required"`
	Difficulty string   `json:"difficulty" binding:"required"`
	Capo       string   `json:"capo"`
	Version    string   `json:"version" binding:"required"`
	Lyrics     string   `json:"lyrics" binding:"required"`
	Youtube    string   `json:"youtube" binding:"required"`
}

func (req *AddNewSongRequest) ToModel() *model.Song {
	return &model.Song{
		Title:      &req.Title,
		Slug:       &req.Slug,
		Artists:    req.Artists,
		Stype:      &req.Stype,
		Difficulty: &req.Difficulty,
		Capo:       &req.Capo,
		Version:    &req.Version,
		Lyrics:     &req.Lyrics,
		Youtube:    &req.Youtube,
	}
}

type FindSongsQuery struct {
	Title string `form:"title"`
	Skip  uint   `form:"skip"`
	Limit uint   `form:"limit,default=10"`
	Sort  string `form:"sort,default='created_at'"`
	Order string `form:"order,default='desc'"`
}
