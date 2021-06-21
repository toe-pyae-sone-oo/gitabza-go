package handler

import (
	"fmt"
	"gitabza-go/common/dateutil"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadArtistPic(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	multi, h, err := c.Request.FormFile("picture")
	if err != nil && err != http.ErrMissingFile {
		log.Printf("error getting picture from form: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}
	defer multi.Close()

	ext := filepath.Ext(h.Filename)[1:]

	filename := fmt.Sprintf("%s-%s.%s", dateutil.GetCurrentDateInStr(), uuid.NewString(), ext)

	file, err := os.Create("public/uploads/" + filename)
	if err != nil {
		log.Printf("error saving artist picture file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}
	defer file.Close()

	b, err := ioutil.ReadAll(multi)
	if err != nil {
		log.Printf("error reading data from multi file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	if _, err := file.Write(b); err != nil {
		log.Printf("error writing data to file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": MsgIntServerErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file": filename})
}
