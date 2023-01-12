package single

type song struct {
	AlbumID      string   `json:"album_id"`
	ArtistIDs    []string `json:"artist_ids"`
	SongGenreID  string   `json:"song_genre_id"`
	SongSourceID string   `json:"source_id"`
	SongID       string   `json:"song_id"`
	SongName     string   `json:"song_name"`
	SongLength   int32    `json:"song_length"`
}
