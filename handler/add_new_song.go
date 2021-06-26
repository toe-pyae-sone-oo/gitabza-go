package handler

import (
	"context"
	"errors"
	"gitabza-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNewSong(c *gin.Context) {
	req, err := getAddNewSongRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := service.NewSongService().Add(context.Background(), req)
	switch {
	case errors.Is(err, service.ErrBadRequest):
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case err != nil:
		log.Printf("error adding new song: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	default:
	}

	c.JSON(http.StatusOK, resp)
}

func getAddNewSongRequest(c *gin.Context) (*service.AddNewSongRequest, error) {
	var req service.AddNewSongRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
