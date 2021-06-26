package handler

import (
	"context"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindSongs(c *gin.Context) {
	query, err := getFindSongsQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	songs, err := service.NewSongService().Find(context.Background(), query)
	if err != nil {
		log.Printf("error finding songs: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func getFindSongsQuery(c *gin.Context) (*service.FindSongsQuery, error) {
	var q service.FindSongsQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		return nil, err
	}
	return &q, nil
}
