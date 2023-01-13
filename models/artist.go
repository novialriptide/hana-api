package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Artist struct {
	ID         primitive.ObjectID `bson:"_id"`
	ArtistID   string             `bson:"artist_id"`
	ArtistName string             `bson:"artist_name"`
}
