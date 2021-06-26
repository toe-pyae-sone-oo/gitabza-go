package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func getArtistUUID(c *gin.Context) (string, error) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return "", errors.New("invalid uuid")
	}
	return uuid, nil
}

func getSongUUID(c *gin.Context) (string, error) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return "", errors.New("invalid uuid")
	}
	return uuid, nil
}
