package service

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/utils"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func (s filmService) GetSpecificFilm(ctx context.Context, request *pb.FilmSpecificRequest) (*pb.FilmSpecificReply, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	key := FormatKeySpecificFilm(request.Id)

	ch := utils.GetSingleFlight().DoChan(key, func() (interface{}, error) {
		var result pb.FilmSpecificReply

		response, err := s.redisRepository.LoadFilmCache(key)

		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(key)
		}()
		if err == nil {
			result = pb.FilmSpecificReply{
				Id:           response.Id,
				Title:        response.Title,
				Resource:     response.Resource,
				State:        response.State,
				Type:         pb.MediaType(pb.MediaType_value[response.Type]),
				Category:     pb.MediaCategory(pb.MediaCategory_value[response.Category]),
				Actors:       response.Actors,
				Directors:    response.Directors,
				Location:     pb.MediaLocation(pb.MediaLocation_value[response.Location]),
				ReleaseYear:  response.ReleaseYear,
				UpdateTime:   timestamppb.New(response.UpdateAt),
				Popularity:   response.Popularity,
				Introduction: response.Introduction,
			}
			if result.Actors == nil {
				result.Actors = []string{}
			}
			if result.Directors == nil {
				result.Directors = []string{}
			}
			return &result, nil
		} else {
			film, err := s.mongodbRepository.GetSpecificFilm(ctx, request)
			if err != nil {
				return nil, err
			}
			result = pb.FilmSpecificReply{
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
			go func() {
				err := s.redisRepository.SaveFilmCache(key, film)
				if err != nil {
					log.Println("error to save cache " + err.Error())
				}
			}()
			return &result, nil
		}
	})

	select {
	case <-ctx.Done():
		return &pb.FilmSpecificReply{}, ctx.Err()
	case ret := <-ch:
		return ret.Val.(*pb.FilmSpecificReply), ret.Err
	}
}
func (s filmService) GetSpecificFilmRoutes(ctx context.Context, request *pb.FilmSpecificRequest) (*pb.FilmSpecificRoutesReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	key := FormatKeySpecificFilmRoutes(request.Id)

	ch := utils.GetSingleFlight().DoChan(key, func() (interface{}, error) {
		var result pb.FilmSpecificRoutesReply

		response, err := s.redisRepository.LoadFilmRoutesCache(key)

		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(key)
		}()
		if err == nil {
			result.Routes = make([]*pb.FilmSpecificRoute, len(response))
			for i, filmRoute := range response {
				result.Routes[i] = &pb.FilmSpecificRoute{Route: filmRoute.Route, Episodes: filmRoute.Episodes}
				if result.Routes[i].Episodes == nil {
					result.Routes[i].Episodes = []string{}
				}
			}
			return &result, nil
		} else {
			filmRoutes, err := s.mongodbRepository.GetSpecificFilmRoutes(ctx, request)
			if err != nil {
				return nil, err
			}
			result.Routes = make([]*pb.FilmSpecificRoute, len(filmRoutes))
			for i, filmRoute := range filmRoutes {
				result.Routes[i] = &pb.FilmSpecificRoute{Route: filmRoute.Route, Episodes: filmRoute.Episodes}
				if result.Routes[i].Episodes == nil {
					result.Routes[i].Episodes = []string{}
				}
			}
			go func() {
				err := s.redisRepository.SaveFilmRoutesCache(key, filmRoutes)
				if err != nil {
					log.Println("error to save cache " + err.Error())
				}
			}()
			return &result, nil
		}
	})

	select {
	case <-ctx.Done():
		return &pb.FilmSpecificRoutesReply{}, ctx.Err()
	case ret := <-ch:
		return ret.Val.(*pb.FilmSpecificRoutesReply), ret.Err
	}
}
