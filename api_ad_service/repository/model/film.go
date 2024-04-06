package model

type Film struct {
	Id          string
	Title       string
	Type        string
	Category    string
	Country     string
	ReleaseYear string

	Resource     string
	State        string
	Popularity   int32
	Introduction string
	UpdateTime   uint32
	Actors       []string
	Directors    []string
}
