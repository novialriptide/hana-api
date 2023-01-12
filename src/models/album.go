package models

type album struct {
	AlbumID   string `json:"album_id"`
	SongIDs   string `json:"song_ids"`
	AlbumName string `json:"album_name"`
}
