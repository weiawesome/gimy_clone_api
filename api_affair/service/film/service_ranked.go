package film

import (
	"api_affair/api/response/film"
	pb "api_affair/proto/film_service"
	"context"
)

func (s *filmService) GetRankedFilms(filmCategory string, offset int32, limit int32) (film.FilmsRanked, error) {
	var result film.FilmsRanked

	client := pb.NewFilmClient(s.c)
	request := pb.FilmRankedRequest{Category: pb.MediaCategory(pb.MediaCategory_value[filmCategory]), Offset: offset, Limit: limit}
	films, err := client.GetRankedFilms(context.Background(), &request)
	if err != nil {
		return result, err
	}

	result.RankedFilms = make([]film.RankedFilm, len(films.RankedFilm))
	for i, filmRanked := range films.RankedFilm {
		result.RankedFilms[i] = film.RankedFilm{
			Id:         filmRanked.Id,
			Title:      filmRanked.Title,
			Popularity: filmRanked.Popularity,
		}
	}
	return result, nil
}
