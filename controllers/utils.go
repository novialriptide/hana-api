package controllers

import (
	"context"
	"hana-api/models"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
)

const uri = "mongodb://localhost:27017"

func ConnectDatabase() *mongo.Client {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return mongoClient
}

func GetConfig() *models.Config {
	var c *models.Config
	file, err := ioutil.ReadFile("config/server.yml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = yaml.Unmarshal(file, c)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return c
}

var mongoClient *mongo.Client = ConnectDatabase()
var config *models.Config = GetConfig()
