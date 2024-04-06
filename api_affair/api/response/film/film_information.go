package film

import "time"

type FilmInformation struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Resource     string    `json:"resource"`
	State        string    `json:"state"`
	Type         string    `json:"type"`
	Category     string    `json:"category"`
	Actors       []string  `json:"actors"`
	Directors    []string  `json:"directors"`
	Location     string    `json:"location"`
	ReleaseYear  uint32    `json:"releaseYear"`
	UpdateTime   time.Time `json:"updateTime"`
	Popularity   int32     `json:"popularity"`
	Introduction string    `json:"introduction"`
}
