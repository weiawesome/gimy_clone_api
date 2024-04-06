package film

import "time"

type FilmsSearch struct {
	SearchFilms []SearchFilm `json:"search_films"`
}
type SearchFilm struct {
	Id           string    `json:"id"`
	State        string    `json:"state"`
	Title        string    `json:"title"`
	Resource     string    `json:"resource"`
	Category     string    `json:"category"`
	Actors       []string  `json:"actors"`
	Directors    []string  `json:"directors"`
	Location     string    `json:"location"`
	Language     string    `json:"language"`
	ReleaseYear  uint32    `json:"release_year"`
	UpdateTime   time.Time `json:"update_time"`
	Introduction string    `json:"introduction"`
}
