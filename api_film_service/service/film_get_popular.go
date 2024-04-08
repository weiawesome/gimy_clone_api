package service

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/utils"
	"context"
	"log"
	"time"
)

func (s filmService) GetTypePopularityFilms(ctx context.Context, request *pb.FilmPopularTypeRequest) (*pb.FilmInformationListReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	key := FormatKeyPopularFilmType(request.Type.String(), request.Offset, request.Limit)

	ch := utils.GetSingleFlight().DoChan(key, func() (interface{}, error) {
		var result pb.FilmInformationListReply

		response, err := s.redisRepository.LoadFilmsCache(key)

		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(key)
		}()
		if err == nil {
			result.FilmInformation = make([]*pb.FilmInformation, len(response))
			for i, film := range response {
				result.FilmInformation[i] = &pb.FilmInformation{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Actors: film.Actors}
				if result.FilmInformation[i].Actors == nil {
					result.FilmInformation[i].Actors = []string{}
				}
			}
			return &result, nil
		} else {
			films, err := s.mongodbRepository.GetPopularFilmsByType(ctx, request)
			if err != nil {
				return nil, err
			}
			result.FilmInformation = make([]*pb.FilmInformation, len(films))
			for i, film := range films {
				result.FilmInformation[i] = &pb.FilmInformation{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Actors: film.Actors}
				if result.FilmInformation[i].Actors == nil {
					result.FilmInformation[i].Actors = []string{}
				}
			}
			go func() {
				err := s.redisRepository.SaveFilmsCache(key, films)
				if err != nil {
					log.Println("error to save cache " + err.Error())
				}
			}()
			return &result, nil
		}
	})

	select {
	case <-ctx.Done():
		return &pb.FilmInformationListReply{}, ctx.Err()
	case ret := <-ch:
		return ret.Val.(*pb.FilmInformationListReply), ret.Err
	}
}
func (s filmService) GetCategoryPopularityFilms(ctx context.Context, request *pb.FilmPopularCategoryRequest) (*pb.FilmInformationListReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	key := FormatKeyPopularFilmCategory(request.Category.String(), request.Offset, request.Limit)

	ch := utils.GetSingleFlight().DoChan(key, func() (interface{}, error) {
		var result pb.FilmInformationListReply

		response, err := s.redisRepository.LoadFilmsCache(key)

		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(key)
		}()
		if err == nil {
			result.FilmInformation = make([]*pb.FilmInformation, len(response))
			for i, film := range response {
				result.FilmInformation[i] = &pb.FilmInformation{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Actors: film.Actors}
				if result.FilmInformation[i].Actors == nil {
					result.FilmInformation[i].Actors = []string{}
				}
			}
			return &result, nil
		} else {
			films, err := s.mongodbRepository.GetPopularFilmsByCategory(ctx, request)
			if err != nil {
				return nil, err
			}
			result.FilmInformation = make([]*pb.FilmInformation, len(films))
			for i, film := range films {
				result.FilmInformation[i] = &pb.FilmInformation{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Actors: film.Actors}
				if result.FilmInformation[i].Actors == nil {
					result.FilmInformation[i].Actors = []string{}
				}
			}
			go func() {
				err := s.redisRepository.SaveFilmsCache(key, films)
				if err != nil {
					log.Println("error to save cache " + err.Error())
				}
			}()
			return &result, nil
		}
	})

	select {
	case <-ctx.Done():
		return &pb.FilmInformationListReply{}, ctx.Err()
	case ret := <-ch:
		return ret.Val.(*pb.FilmInformationListReply), ret.Err
	}
}
