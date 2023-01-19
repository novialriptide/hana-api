package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSongs(ginContext *gin.Context) {
	collection := mongoClient.Database("hana-db").Collection("songs")

	collection.Find(context.TODO(), bson.D{})

}
