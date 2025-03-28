package jellyfin

type session []struct {
	PlayState struct {
		IsPaused   bool   `json:"IsPaused"`
		PlayMethod string `json:"PlayMethod"`
	} `json:"PlayState,omitempty"`
	UserName string `json:"UserName"`
	IsActive bool   `json:"IsActive"`
}

type mediaCounts struct {
	AlbumCount      int `json:"AlbumCount"`
	ArtistCount     int `json:"ArtistCount"`
	BookCount       int `json:"BookCount"`
	BoxSetCount     int `json:"BoxSetCount"`
	EpisodeCount    int `json:"EpisodeCount"`
	ItemCount       int `json:"ItemCount"`
	MovieCount      int `json:"MovieCount"`
	MusicVideoCount int `json:"MusicVideoCount"`
	ProgramCount    int `json:"ProgramCount"`
	SeriesCount     int `json:"SeriesCount"`
	SongCount       int `json:"SongCount"`
	TrailerCount    int `json:"TrailerCount"`
}
