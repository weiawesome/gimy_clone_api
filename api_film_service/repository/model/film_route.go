package model

type FilmRoute struct {
	Id       string   `bson:"id"`
	Route    string   `bson:"route"`
	Episodes []string `bson:"episodes"`
}
