package request

import "time"

type AdInformation struct {
	Bucket        string    `json:"bucket"`
	Id            string    `json:"id"`
	FileExtension string    `json:"file_extension"`
	ExpireTime    time.Time `json:"expire_time"`
}
