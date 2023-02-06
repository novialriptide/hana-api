package mongo_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	ID             primitive.ObjectID `bson:"_id"`
	SongID         string             `bson:"song_id"`
	AlbumID        string             `bson:"album_id"`
	ArtistIDs      []string           `bson:"artist_ids"`
	SongGenreID    string             `bson:"song_genre_id"`
	UploaderUserID string             `bson:"uploader_user_id"`
	SongFileSource string             `bson:"song_file_source"`
	SongFileID     primitive.ObjectID `bson:"song_file_id"`
	SongName       string             `bson:"song_name"`
}
