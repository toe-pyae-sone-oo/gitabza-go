package handler

import (
	"context"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindArtists(c *gin.Context) {
	query, err := getFindArtistsQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	artists, err := service.NewArtistService().Find(context.Background(), query)
	if err != nil {
		log.Printf("error finding artists: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	c.JSON(http.StatusOK, artists)
}

func getFindArtistsQuery(c *gin.Context) (*service.FindArtistsQuery, error) {
	var query service.FindArtistsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		return nil, err
	}
	return &query, nil
}
