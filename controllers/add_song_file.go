package controllers

import (
	"hana-api/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddSongFile(ginContext *gin.Context) {
	database := mongoClient.Database("hana-db")
	// songID := ginContext.Param("song_id")

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

	data, err := ioutil.ReadAll(openedFile)
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
	uploadOpts := options.GridFSUpload().SetChunkSizeBytes(200000)
	uploadStream, err := bucket.OpenUploadStream("file.txt", uploadOpts)
	if err != nil {
		panic(err)
	}

	if _, err = uploadStream.Write(data); err != nil {
		panic(err)
	}
	ginContext.IndentedJSON(http.StatusOK, models.Result{
		IsSuccessful: true,
		Message:      "Successfully uploaded",
	})
}
