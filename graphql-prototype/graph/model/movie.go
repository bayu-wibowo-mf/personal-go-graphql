package model

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ReleaseDate string `json:"releaseDate"`
}

type NewMovie struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
