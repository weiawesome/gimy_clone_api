package film

type FilmsRanked struct {
	RankedFilms []RankedFilm `json:"ranked_films"`
}
type RankedFilm struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Popularity int32  `json:"popularity"`
}
