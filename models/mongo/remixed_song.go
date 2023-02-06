package mongo_models

type RemixedSong struct {
	Song
	RemixArtistID  string `bson:"remix_artist_id"`
	OriginalSongID string `bson:"original_song_id"`
}
