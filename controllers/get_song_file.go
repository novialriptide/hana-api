package controllers

import (
	"bytes"
	"context"
	"hana-api/models"
	mongo_models "hana-api/models/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSongFile(ginContext *gin.Context) {
	database := mongoClient.Database("hana-db")
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

	bucket, err := gridfs.NewBucket(database, options.GridFSBucket().SetName("musicdata"))
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}

	fileBuffer := bytes.NewBuffer(nil)
	if _, err := bucket.DownloadToStream(song.SongFileID, fileBuffer); err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}

	filter = bson.D{{Key: "_id", Value: song.SongFileID}}

	var metadata gridfs.File
	metadataResult := bucket.GetFilesCollection().FindOne(context.TODO(), filter)
	metadataResult.Decode(&metadata)

	ginContext.Header("Content-Disposition", "attachment; filename="+metadata.Name)
	_, err = ginContext.Writer.Write(fileBuffer.Bytes())
	if err != nil {
		ginContext.IndentedJSON(http.StatusInternalServerError, models.Result{
			IsSuccessful: false,
			Message:      err.Error(),
		})
		panic(err)
	}

	ginContext.IndentedJSON(http.StatusOK, models.Result{
		IsSuccessful: true,
		Message:      "Download successful",
	})
}
