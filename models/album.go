package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// for getting song ids, just search for songs that have the albumid set to the actual album

type Album struct {
	ID        primitive.ObjectID `bson:"_id"`
	AlbumID   string             `bson:"album_id"`
	AlbumName string             `bson:"album_name"`
}
