package service

import (
	pb "api_film_service/proto/film_service"
	"context"
)

func (s filmService) AddPopularity(ctx context.Context, request *pb.FilmSpecificRequest) error {
	return s.mongodbRepository.AddPopularity(ctx, request)
}
func (s filmService) AddFilmToSearchEngine(ctx context.Context, request *pb.FilmSpecificRequest) error {
	film, err := s.mongodbRepository.GetSpecificFilm(ctx, request)
	if err != nil {
		return err
	}
	return s.elasticsearchRepository.FilmSave(ctx, film)
}
