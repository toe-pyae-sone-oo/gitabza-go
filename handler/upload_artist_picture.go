package handler

import (
	"fmt"
	"gitabza-go/common/artistutil"
	"gitabza-go/common/dateutil"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadArtistPic(c *gin.Context) {
	multi, filename, err := getMultiFile(c, "picture")
	switch {
	case err == http.ErrMissingFile:
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing file in picture field"})
		return
	case err != nil:
		log.Printf("error getting artist picture file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	default:
	}
	defer multi.Close()

	ext := filepath.Ext(filename)[1:]
	if !artistutil.IsFileSupported(ext) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file not supported"})
		return
	}

	filename = fmt.Sprintf("%s-%s.%s", dateutil.GetCurrentDateInStr(), uuid.NewString(), ext)
	if err := saveFile("public/uploads/"+filename, multi); err != nil {
		log.Printf("error saving artist picture file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file": filename})
}

func getMultiFile(c *gin.Context, field string) (multipart.File, string, error) {
	if err := c.Request.ParseForm(); err != nil {
		return nil, "", err
	}

	multi, h, err := c.Request.FormFile(field)
	if err != nil {
		return nil, "", err
	}

	return multi, h.Filename, nil
}

func saveFile(filename string, reader io.Reader) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	if _, err := file.Write(b); err != nil {
		return err
	}

	return nil
}
