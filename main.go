package main

import (
	"hana-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{AlbumID: "0", SongIDs: []string{"52", "14"}, AlbumName: "Lmao"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/get-albums", getAlbums)

	router.Run("localhost:25565")
}
