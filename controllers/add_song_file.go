package controllers

import (
	"context"
	"hana-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddSongFile(ginContext *gin.Context) {
	collection := mongoClient.Database("hana-db").Collection("songfiles")
	songID := ginContext.Query("song_id")
	file, fileErr := ginContext.FormFile(songID)
	if fileErr != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      fileErr.Error(),
		})
		return
	}

	ginContext.SaveUploadedFile(file, config.SongFilesPath)

	s := models.SongFile{
		ID:       primitive.NewObjectID(),
		SongID:   songID,
		FilePath: config.SongFilesPath,
	}

	_, mongoErr := collection.InsertOne(context.TODO(), s)
	if mongoErr != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      mongoErr.Error(),
		})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, models.Result{
		IsSuccessful: true,
		Message:      "Uploaded a new song file",
	})
}
