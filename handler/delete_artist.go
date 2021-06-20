package handler

import (
	"context"
	"fmt"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func getArtistUUID(c *gin.Context) (string, error) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return "", errors.New("invalid uuid")
	}
	return uuid, nil
}
