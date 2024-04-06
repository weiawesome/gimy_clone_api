package request

type FilmInformation struct {
	Id            string `json:"id"`
	Route         string `json:"route"`
	Episode       string `json:"episode"`
	FileExtension string `json:"file_extension"`
	State         string `json:"state"`
}
