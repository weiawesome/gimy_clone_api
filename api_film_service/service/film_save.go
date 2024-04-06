package service

import (
	pb "api_film_service/proto/film_service"
	"context"
	"log"
)

func (s filmService) SaveFilm(ctx context.Context, request *pb.FilmSaveRequest) error {
	if request.Actors == nil {
		request.Actors = []string{}
	}
	if request.Directors == nil {
		request.Directors = []string{}
	}
	go func() {
		err := s.AddFilmToSearchEngine(ctx, &pb.FilmSpecificRequest{Id: request.Id})
		if err != nil {
			log.Println("error to add film id (" + request.Id + ") to search engine with " + err.Error())
		}
	}()
	return s.mongodbRepository.SaveFilm(ctx, request)
}
func (s filmService) SaveFilmEpisode(ctx context.Context, request *pb.FilmSaveEpisodeRequest) error {
	return s.mongodbRepository.SaveFilmEpisode(ctx, request)
}
