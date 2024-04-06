package utils

import "time"

type FilmInformation struct {
	Id            string `json:"id"`
	Route         string `json:"route"`
	Episode       string `json:"episode"`
	FileExtension string `json:"file_extension"`
	State         string `json:"state"`
}
type AdInformation struct {
	Bucket        string    `json:"bucket"`
	Id            string    `json:"id"`
	FileExtension string    `json:"file_extension"`
	ExpireTime    time.Time `json:"expire_time"`
}
