package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SongFile struct {
	// Use MongoDB's ObjectID to get file upload date
	ID       primitive.ObjectID `bson:"_id"`
	SongID   string             `bson:"song_id"`
	FilePath string             `bson:"file_path"`
}
