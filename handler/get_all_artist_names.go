package handler

import (
	"context"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllArtistNames(c *gin.Context) {
	names, err := service.NewArtistService().GetAllNames(context.Background())
	if err != nil {
		log.Printf("error getting artist names: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}
	c.JSON(http.StatusOK, names)
}
