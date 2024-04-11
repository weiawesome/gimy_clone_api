package api

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/service"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedFilmServer
}

func (s *Server) GetBasicFilms(ctx context.Context, in *pb.FilmBasicRequest) (*pb.FilmInformationListReply, error) {
	return service.GetService().GetBasicFilms(ctx, in)
}

func (s *Server) GetPopularTypeFilms(ctx context.Context, in *pb.FilmPopularTypeRequest) (*pb.FilmInformationListReply, error) {
	return service.GetService().GetTypePopularityFilms(ctx, in)
}

func (s *Server) GetPopularCategoryFilms(ctx context.Context, in *pb.FilmPopularCategoryRequest) (*pb.FilmInformationListReply, error) {
	return service.GetService().GetCategoryPopularityFilms(ctx, in)
}

func (s *Server) GetRankedFilms(ctx context.Context, in *pb.FilmRankedRequest) (*pb.FilmRankedReply, error) {
	return service.GetService().GetRankedFilms(ctx, in)
}
func (s *Server) GetFilterFilms(ctx context.Context, in *pb.FilmFilterRequest) (*pb.FilmInformationListReply, error) {
	return service.GetService().GetFilterFilms(ctx, in)
}
func (s *Server) GetSpecificFilm(ctx context.Context, in *pb.FilmSpecificRequest) (*pb.FilmSpecificReply, error) {
	return service.GetService().GetSpecificFilm(ctx, in)
}
func (s *Server) GetSpecificFilmRoutes(ctx context.Context, in *pb.FilmSpecificRequest) (*pb.FilmSpecificRoutesReply, error) {
	return service.GetService().GetSpecificFilmRoutes(ctx, in)
}
func (s *Server) GetSearchFilms(ctx context.Context, in *pb.FilmSearchRequest) (*pb.FilmSearchReply, error) {
	return service.GetService().GetSearchFilms(ctx, in)
}

func (s *Server) SaveFilm(ctx context.Context, in *pb.FilmSaveRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.GetService().SaveFilm(ctx, in)
}

func (s *Server) SaveFilmEpisode(ctx context.Context, in *pb.FilmSaveEpisodeRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.GetService().SaveFilmEpisode(ctx, in)
}
func (s *Server) DeleteFilmEpisode(ctx context.Context, in *pb.FilmSaveEpisodeRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.GetService().DeleteFilmEpisode(ctx, in)
}
func (s *Server) DeleteFilm(ctx context.Context, in *pb.FilmSpecificRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.GetService().DeleteFilm(ctx, in)
}
func (s *Server) AddFilmToSearchEngine(ctx context.Context, in *pb.FilmSpecificRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.GetService().AddFilmToSearchEngine(ctx, in)
}
func (s *Server) AddFilmPopularity(ctx context.Context, in *pb.FilmSpecificRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.GetService().AddPopularity(ctx, in)
}
