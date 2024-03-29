package handler

import "github.com/gin-gonic/gin"

func HandleRoutes(r *gin.Engine) {
	artistGroup := r.Group("/artists")
	{
		artistGroup.POST("/upload/pic", UploadArtistPic)
		artistGroup.POST("/", AddNewArtist)
		artistGroup.GET("/", FindArtists)
		artistGroup.DELETE("/:uuid", DeleteArtist)
		artistGroup.GET("/names", GetAllArtistNames)
		artistGroup.GET("/slug/:slug", GetArtistBySlug)
		artistGroup.GET("/:uuid", GetArtistByUUID)
		artistGroup.GET("/:uuid/songs", GetArtistSongs)
		artistGroup.PUT("/:uuid", UpdateArtist)
	}

	songGroup := r.Group("/songs")
	{
		songGroup.POST("/", AddNewSong)
		songGroup.GET("/", FindSongs)
		songGroup.DELETE("/:uuid", DeleteSong)
		songGroup.GET("/slug/:artist/:song", GetSongBySlug)
	}
}
