package mongo_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RemixedSong struct {
	ID primitive.ObjectID `bson:"_id"`
	Song
	OriginalSongID string `bson:"original_song_id"`
}
