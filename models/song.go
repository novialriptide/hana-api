package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	ID           primitive.ObjectID `bson:"_id"`
	SongID       string             `bson:"song_id"`
	AlbumID      string             `bson:"album_id"`
	ArtistIDs    []string           `bson:"artist_ids"`
	SongGenreID  string             `bson:"song_genre_id"`
	SongSourceID string             `bson:"source_id"`
	SongName     string             `bson:"song_name"`
	SongLength   int32              `bson:"song_length"`
}
