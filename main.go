package main

import (
	"context"
	"fmt"
	"hana-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://localhost:27017"

var albums = []models.Album{
	{AlbumID: "0", SongIDs: []string{"52", "14"}, AlbumName: "Lmao"},
}

var mongoClient *mongo.Client

func connectDatabase() *mongo.Client {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return mongoClient
}

func getSongByID(ginContext *gin.Context) {
	songID := ginContext.Param("song_id")
	filter := bson.D{{Key: "song_id", Value: songID}}
	collection := mongoClient.Database("hana-db").Collection("songs")

	var result models.Song

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	ginContext.IndentedJSON(http.StatusOK, result)
}

func getAlbumByID(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func main() {
	mongoClient = connectDatabase()

	router := gin.Default()
	router.GET("/songs/:song_id", getSongByID)
	router.GET("/albums", getAlbumByID)

	router.Run("localhost:25565")
}
