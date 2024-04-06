package model

import "time"

type Advertisement struct {
	ExpireAt time.Time `bson:"expire_at"`
	Bucket   string    `bson:"bucket"`
	File     string    `bson:"file"`
}
