package handler

import (
	"context"
	"fmt"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSong(c *gin.Context) {
	uuid, err := getSongUUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := service.NewSongService().Delete(context.Background(), uuid); err != nil {
		log.Printf("error deleting song: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("song with id %s successfully deleted", uuid),
	})
}
