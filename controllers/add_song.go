package controllers

import (
	"context"
	"hana-api/models"
	mongo_models "hana-api/models/mongo"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddSong(ginContext *gin.Context) {
	collection := mongoClient.Database("hana-db").Collection("songs")

	s := mongo_models.Song{
		ID:             primitive.NewObjectID(),
		SongID:         uuid.New().String(),
		AlbumID:        ginContext.Query("album_id"),
		ArtistIDs:      strings.Split(ginContext.Query("artist_ids"), ","),
		SongGenreID:    ginContext.Query("song_genre_id"),
		SongFileSource: ginContext.Query("song_file_source"),
		SongName:       ginContext.Query("song_name"),
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
		Message:      s.SongID,
	})
}
