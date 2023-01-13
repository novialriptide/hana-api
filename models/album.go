package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID        primitive.ObjectID `bson:"_id"`
	AlbumID   string             `bson:"album_id"`
	SongIDs   []string           `bson:"song_ids"`
	AlbumName string             `bson:"album_name"`
}
