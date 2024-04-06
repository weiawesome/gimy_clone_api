package service

import (
	pb "api_film_service/proto/film_service"
	"context"
)

func (s filmService) GetFilterFilms(ctx context.Context, request *pb.FilmFilterRequest) (*pb.FilmInformationListReply, error) {
	films, err := s.mongodbRepository.GetFilterFilms(ctx, request)
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
func (s filmService) GetRankedFilms(ctx context.Context, request *pb.FilmRankedRequest) (*pb.FilmRankedReply, error) {
	films, err := s.mongodbRepository.GetRankedFilms(ctx, request)
	if err != nil {
		return nil, err
	}
	var result pb.FilmRankedReply
	result.RankedFilm = make([]*pb.FilmRanked, len(films))
	for i, film := range films {
		result.RankedFilm[i] = &pb.FilmRanked{Id: film.Id, Title: film.Title, Popularity: film.Popularity}
	}
	return &result, nil
}
