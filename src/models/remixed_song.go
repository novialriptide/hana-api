package models

type remixedSong struct {
	song
	OriginalSongID string `json:"original_song_id"`
}
