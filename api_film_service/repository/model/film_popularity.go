package model

import "time"

type FilmPopularity struct {
	Id         string    `bson:"id"`
	CreatedAt  time.Time `bson:"created_at"`
	popularity int64     `bson:"popularity"`
}
