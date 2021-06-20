package repository

import "errors"

var ErrInvalidParm = errors.New("invalid param")

const (
	DbName     = "gitabza"
	ColSongs   = "songs"
	ColArtists = "artists"
)
