package main

import (
	"hana-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Get an existing artist
	router.GET("/artists/:artist_id", controllers.GetArtistByID)

	// Add a new artist profile
	router.POST("/artists", controllers.AddArtist)

	// Get an existing album
	router.GET("/albums/:album_id", controllers.GetAlbumByID)

	// Add a new album
	router.POST("/albums", controllers.AddAlbum)

	// Get x number of songs
	router.GET("/songs", controllers.GetSongs)

	// Get an existing song
	router.GET("/songs/:song_id", controllers.GetSongByID)

	// Add a new song
	router.POST("/songs", controllers.AddSong)

	// Add a new source to an existing song
	// TODO: router.POST("/songs/:song_id/file")

	router.Run("localhost:25565")
}
