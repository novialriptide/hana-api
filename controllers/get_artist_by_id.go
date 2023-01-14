package controllers

import (
	"context"
	"hana-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetArtistByID(ginContext *gin.Context) {
	artistID := ginContext.Param("artist_id")
	filter := bson.D{{Key: "artist_id", Value: artistID}}
	collection := mongoClient.Database("hana-db").Collection("artists")

	var result models.Artist

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
