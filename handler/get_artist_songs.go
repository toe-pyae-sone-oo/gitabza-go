package handler

import (
	"context"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArtistSongs(c *gin.Context) {
	uuid, err := getArtistUUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := service.NewSongService().FindByArtist(context.Background(), uuid)
	if err != nil {
		log.Printf("error finding artist songs: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	c.JSON(http.StatusOK, resp)
}
