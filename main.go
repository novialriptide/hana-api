package main

import (
	"hana-server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Get x number of songs
	router.GET("/songs", controllers.GetSongs)

	// Get an existing song
	router.GET("/songs/:song_id", controllers.GetSongByID)

	// Add a new song
	router.POST("/songs", controllers.AddSong)

	// Add a new source to an existing song
	// TODO: router.POST("/songs/:song_id/file")

	// Add a new album
	router.POST("/albums", controllers.AddAlbum)

	router.Run("localhost:25565")
}
