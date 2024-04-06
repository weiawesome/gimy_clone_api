package service

import (
	pb "api_film_service/proto/film_service"
	"context"
)

func (s filmService) GetTypePopularityFilms(ctx context.Context, request *pb.FilmPopularTypeRequest) (*pb.FilmInformationListReply, error) {
	films, err := s.mongodbRepository.GetPopularFilmsByType(ctx, request)
	if err != nil {
		return nil, err
	}
	var result pb.FilmInformationListReply
	result.FilmInformation = make([]*pb.FilmInformation, len(films))
	for i, film := range films {
		result.FilmInformation[i] = &pb.FilmInformation{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Actors: film.Actors}
		if result.FilmInformation[i].Actors == nil {
			result.FilmInformation[i].Actors = []string{}
		}
	}
	return &result, nil
}
func (s filmService) GetCategoryPopularityFilms(ctx context.Context, request *pb.FilmPopularCategoryRequest) (*pb.FilmInformationListReply, error) {
	films, err := s.mongodbRepository.GetPopularFilmsByCategory(ctx, request)
	if err != nil {
		return nil, err
	}
	var result pb.FilmInformationListReply
	result.FilmInformation = make([]*pb.FilmInformation, len(films))
	for i, film := range films {
		result.FilmInformation[i] = &pb.FilmInformation{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Actors: film.Actors}
		if result.FilmInformation[i].Actors == nil {
			result.FilmInformation[i].Actors = []string{}
		}
	}
	return &result, nil
}
