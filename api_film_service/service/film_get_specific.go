package service

import (
	pb "api_film_service/proto/film_service"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s filmService) GetSpecificFilm(ctx context.Context, request *pb.FilmSpecificRequest) (*pb.FilmSpecificReply, error) {
	film, err := s.mongodbRepository.GetSpecificFilm(ctx, request)
	if err != nil {
		return nil, err
	}
	result := pb.FilmSpecificReply{
		Id:           film.Id,
		Title:        film.Title,
		Resource:     film.Resource,
		State:        film.State,
		Type:         pb.MediaType(pb.MediaType_value[film.Type]),
		Category:     pb.MediaCategory(pb.MediaCategory_value[film.Category]),
		Actors:       film.Actors,
		Directors:    film.Directors,
		Location:     pb.MediaLocation(pb.MediaLocation_value[film.Location]),
		ReleaseYear:  film.ReleaseYear,
		UpdateTime:   timestamppb.New(film.UpdateAt),
		Popularity:   film.Popularity,
		Introduction: film.Introduction,
	}
	if result.Actors == nil {
		result.Actors = []string{}
	}
	if result.Directors == nil {
		result.Directors = []string{}
	}
	return &result, nil
}
func (s filmService) GetSpecificFilmRoutes(ctx context.Context, request *pb.FilmSpecificRequest) (*pb.FilmSpecificRoutesReply, error) {
	films, err := s.mongodbRepository.GetSpecificFilmRoutes(ctx, request)
	if err != nil {
		return nil, err
	}
	var result pb.FilmSpecificRoutesReply
	result.Routes = make([]*pb.FilmSpecificRoute, len(films))
	for i, film := range films {
		result.Routes[i] = &pb.FilmSpecificRoute{Route: film.Route, Episodes: film.Episodes}
		if result.Routes[i].Episodes == nil {
			result.Routes[i].Episodes = []string{}
		}
	}
	return &result, nil
}
