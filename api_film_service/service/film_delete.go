package service

import (
	pb "api_film_service/proto/film_service"
	"context"
)

func (s filmService) DeleteFilmEpisode(ctx context.Context, request *pb.FilmSaveEpisodeRequest) error {
	return s.mongodbRepository.DeleteFilmEpisode(ctx, request)
}

func (s filmService) DeleteFilm(ctx context.Context, request *pb.FilmSpecificRequest) error {
	return s.mongodbRepository.DeleteFilm(ctx, request)
}
