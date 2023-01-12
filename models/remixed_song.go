package models

type RemixedSong struct {
	Song
	OriginalSongID string `json:"original_song_id"`
}
