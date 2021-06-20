package handler

import (
	"context"
	"errors"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArtistBySlug(c *gin.Context) {
	slug, err := getArtistSlug(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := service.NewArtistService().FindBySlug(context.Background(), slug)
	if err != nil {
		log.Printf("error getting artist by slug: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func getArtistSlug(c *gin.Context) (string, error) {
	slug := c.Param("slug")
	if slug == "" {
		return "", errors.New("invalid slug")
	}
	return slug, nil
}
