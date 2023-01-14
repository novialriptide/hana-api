package controllers

import (
	"context"
	"hana-server/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddSong(ginContext *gin.Context) {
	collection := mongoClient.Database("hana-db").Collection("songs")

	s := models.Song{
		ID:          primitive.NewObjectID(),
		SongID:      uuid.New().String(),
		AlbumID:     ginContext.Query("album_id"),
		ArtistIDs:   strings.Split(ginContext.Query("artist_ids"), ","),
		SongGenreID: ginContext.Query("song_genre_id"),
		SongSource:  ginContext.Query("song_source"),
		SongName:    ginContext.Query("song_name"),
	}

	_, err := collection.InsertOne(context.TODO(), s)

	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, models.Result{
		IsSuccessful: true,
		Message:      "Added a new song",
	})
}
