package handler

import (
	"context"
	"errors"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSongBySlug(c *gin.Context) {
	songSlug, artistSlug, err := getSongAndArtistSlug(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	s := service.NewSongService()
	resp, err := s.FindBySlug(ctx, songSlug)
	switch {
	case errors.Is(err, service.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	case err != nil:
		log.Printf("error finding song by slug: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	related, err := s.IsSongRelatedToArtist(ctx, resp, artistSlug)
	switch {
	case errors.Is(err, service.ErrBadRequest):
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case err != nil:
		log.Printf("error finding song relation to artist: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	case !related:
		c.JSON(http.StatusNotFound, gin.H{
			"message": "song not found",
		})
		return
	default:
	}

	c.JSON(http.StatusOK, resp)
}

func getSongAndArtistSlug(c *gin.Context) (songSlug, artistSlug string, err error) {
	songSlug, artistSlug = c.Param("song"), c.Param("artist")
	if songSlug == "" || artistSlug == "" {
		err = errors.New("invalid song and artist slugs")
		return
	}
	return
}
