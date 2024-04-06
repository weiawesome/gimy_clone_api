package service

import (
	pb "api_film_service/proto/film_service"
	"context"
)

func (s filmService) GetBasicFilms(ctx context.Context, request *pb.FilmBasicRequest) (*pb.FilmInformationListReply, error) {
	films, err := s.mongodbRepository.GetBasicFilms(ctx, request)
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
