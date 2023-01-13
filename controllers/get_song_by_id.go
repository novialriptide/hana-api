package controllers

import (
	"context"
	"hana-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSongByID(ginContext *gin.Context) {
	songID := ginContext.Param("song_id")
	filter := bson.D{{Key: "song_id", Value: songID}}
	collection := mongoClient.Database("hana-db").Collection("songs")

	var result models.Song

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
	}

	ginContext.IndentedJSON(http.StatusOK, result)
}
