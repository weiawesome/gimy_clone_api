package service

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/model"
	"context"
	"log"
	"time"
)

func (s filmService) SaveFilm(ctx context.Context, request *pb.FilmSaveRequest) error {
	if request.Actors == nil {
		request.Actors = []string{}
	}
	if request.Directors == nil {
		request.Directors = []string{}
	}

	err := s.mongodbRepository.SaveFilm(ctx, request)
	if err != nil {
		return err
	}
	go func() {
		film := model.Film{
			Id:           request.Id,
			Title:        request.Title,
			Type:         request.Type.String(),
			Category:     request.Category.String(),
			Location:     request.Location.String(),
			Language:     request.Language,
			ReleaseYear:  request.ReleaseYear,
			Resource:     request.Resource,
			State:        request.State,
			Popularity:   0,
			Introduction: request.Introduction,
			Actors:       request.Actors,
			Directors:    request.Directors,
			UpdateAt:     time.Now().UTC(),
			CreatedAt:    time.Now().UTC(),
		}
		err := s.elasticsearchRepository.FilmSave(ctx, film)
		if err != nil {
			log.Println("error to add film id (" + request.Id + ") to search engine with " + err.Error())
		}
	}()
	return nil
}
func (s filmService) SaveFilmEpisode(ctx context.Context, request *pb.FilmSaveEpisodeRequest) error {
	return s.mongodbRepository.SaveFilmEpisode(ctx, request)
}
