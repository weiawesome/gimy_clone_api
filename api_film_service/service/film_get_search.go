package service

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/model"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s filmService) GetSearchFilms(ctx context.Context, request *pb.FilmSearchRequest) (*pb.FilmSearchReply, error) {
	var films []model.Film
	var err error
	switch request.SearchType {
	case pb.SearchType_TITLE:
		films, err = s.elasticsearchRepository.SearchFilmsGet(ctx, request)
		if err != nil {
			return nil, err
		}
	case pb.SearchType_CELEBRITY:
		films, err = s.mongodbRepository.GetSearchFilms(ctx, request)
		if err != nil {
			return nil, err
		}
	}
	var result pb.FilmSearchReply
	result.Results = make([]*pb.FilmSearchResult, len(films))
	for i, film := range films {
		result.Results[i] = &pb.FilmSearchResult{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Category: pb.MediaCategory(pb.MediaCategory_value[film.Category]), Actors: film.Actors, Director: film.Directors, Location: pb.MediaLocation(pb.MediaLocation_value[film.Location]), Language: film.Language, UpdateTime: timestamppb.New(film.UpdateAt), Introduction: film.Introduction, ReleaseYear: film.ReleaseYear}
		if result.Results[i].Actors == nil {
			result.Results[i].Actors = []string{}
		}
		if result.Results[i].Director == nil {
			result.Results[i].Director = []string{}
		}
	}
	return &result, nil
}
