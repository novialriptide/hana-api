package controllers

import (
	"context"
	"hana-api/models"
	mongo_models "hana-api/models/mongo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSongs(ginContext *gin.Context) {
	collection := mongoClient.Database("hana-db").Collection("songs")

	var songs []mongo_models.Song

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      "error retriving songs",
		})
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var result mongo_models.Song
		if err := cursor.Decode(&result); err != nil {
			log.Println(err)
		}
		songs = append(songs, result)
	}

	ginContext.IndentedJSON(http.StatusOK, songs)
}
