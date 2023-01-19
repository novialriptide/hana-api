package controllers

import (
	"context"
	"hana-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddArtist(ginContext *gin.Context) {
	collection := mongoClient.Database("hana-db").Collection("artists")

	a := models.Artist{
		ID:         primitive.NewObjectID(),
		ArtistID:   uuid.New().String(),
		ArtistName: ginContext.Query("artist_name"),
	}

	_, err := collection.InsertOne(context.TODO(), a)

	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		return
	}

	ginContext.IndentedJSON(http.StatusOK, models.Result{
		IsSuccessful: true,
		Message:      "Added a new artist",
	})
}
