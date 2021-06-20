package handler

import (
	"context"
	"errors"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateArtist(c *gin.Context) {
	req, err := getUpdateArtistRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := service.NewArtistService().Update(context.Background(), req)
	switch {
	case errors.Is(err, service.ErrBadRequest):
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case err != nil:
		log.Printf("error updating artist: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	default:
	}

	c.JSON(http.StatusOK, resp)
}

func getUpdateArtistRequest(c *gin.Context) (*service.UpdateArtistRequest, error) {
	var req service.UpdateArtistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	uuid := c.Param("uuid")
	if uuid == "" {
		return nil, errors.New("invalid uuid")
	}

	req.UUID = uuid

	return &req, nil
}
