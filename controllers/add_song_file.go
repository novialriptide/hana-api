package controllers

import (
	"context"
	"hana-api/models"
	mongo_models "hana-api/models/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddSongFile(ginContext *gin.Context) {
	database := mongoClient.Database("hana-db")

	// Check if song data exists and fetch variables
	songID := ginContext.Param("song_id")
	filter := bson.D{{Key: "song_id", Value: songID}}
	collection := mongoClient.Database("hana-db").Collection("songs")

	var song mongo_models.Song
	err := collection.FindOne(context.TODO(), filter).Decode(&song)
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      "song data does not exist",
		})
		return
	}

	// Handle getting music data from RESTful API
	fileForm, err := ginContext.FormFile("file")
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}

	openedFile, err := fileForm.Open()
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}

	// Handle adding the music data to MongoDB using GridFS
	bucket, err := gridfs.NewBucket(database, options.GridFSBucket().SetName("musicdata"))
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}
	objectID, err := bucket.UploadFromStream(fileForm.Filename, openedFile)
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}

	// Assign file to song ID
	song.SongSourceID = objectID
	_, err = collection.ReplaceOne(context.TODO(), filter, song)
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}

	// Output
	ginContext.IndentedJSON(http.StatusOK, models.Result{
		IsSuccessful: true,
		Message:      "Successfully uploaded",
	})
}
