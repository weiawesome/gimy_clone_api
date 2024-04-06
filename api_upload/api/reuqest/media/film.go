package media

type FilmInformation struct {
	Title        string   `json:"title"`
	State        string   `json:"state"`
	Type         string   `json:"type"`
	Category     string   `json:"category"`
	Actors       []string `json:"actors"`
	Directors    []string `json:"directors"`
	Location     string   `json:"location"`
	ReleaseYear  uint32   `json:"releaseYear"`
	Introduction string   `json:"introduction"`
	Language     string   `json:"language"`
}
