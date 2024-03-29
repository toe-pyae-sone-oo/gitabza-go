package handler

import (
	"context"
	"fmt"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteArtist(c *gin.Context) {
	uuid, err := getArtistUUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := service.NewArtistService().Delete(context.Background(), uuid); err != nil {
		log.Printf("error deleting artist: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{"message": fmt.Sprintf("artist with id %s successfully deleted", uuid)})
}
