package controllers

import (
	"context"
	"hana-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAlbumByID(ginContext *gin.Context) {
	albumID := ginContext.Param("album_id")
	filter := bson.D{{Key: "album_id", Value: albumID}}
	collection := mongoClient.Database("hana-db").Collection("albums")

	var result models.Album

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, result)
}
