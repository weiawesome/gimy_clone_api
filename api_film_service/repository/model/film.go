package model

import "time"

type Film struct {
	Id           string   `bson:"id"`
	Title        string   `bson:"title"`
	Type         string   `bson:"type"`
	Category     string   `bson:"category"`
	Location     string   `bson:"location"`
	Language     string   `bson:"language"`
	ReleaseYear  uint32   `bson:"release_year"`
	Resource     string   `bson:"resource"`
	State        string   `bson:"state"`
	Popularity   int32    `bson:"popularity"`
	Introduction string   `bson:"introduction"`
	Actors       []string `bson:"actors"`
	Directors    []string `bson:"directors"`

	UpdateAt  time.Time `bson:"update_at"`
	CreatedAt time.Time `bson:"created_at"`
}
