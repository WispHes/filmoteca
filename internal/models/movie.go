package models

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
	Actors      []int   `json:"actors"`
}
