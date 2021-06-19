package handler

import (
	"context"
	"log"

	"gitabza-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AddNewArtist(c *gin.Context) {
	req, err := getAddNewArtistRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := service.NewArtistService().Add(context.Background(), req)
	switch {
	case errors.Is(err, service.ErrBadRequest):
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case err != nil:
		log.Printf("error adding new artist: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	default:
	}

	c.JSON(http.StatusOK, resp)
}

func getAddNewArtistRequest(c *gin.Context) (*service.AddNewArtistRequest, error) {
	var req service.AddNewArtistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
