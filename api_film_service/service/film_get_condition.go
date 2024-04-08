package service

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/utils"
	"context"
	"log"
	"time"
)

func (s filmService) GetFilterFilms(ctx context.Context, request *pb.FilmFilterRequest) (*pb.FilmInformationListReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	key := FormatKeyFilter(request.Type.String(), request.Category.String(), request.ReleaseYear, request.OrderType.String(), request.Offset, request.Limit)

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
			films, err := s.mongodbRepository.GetFilterFilms(ctx, request)
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
func (s filmService) GetRankedFilms(ctx context.Context, request *pb.FilmRankedRequest) (*pb.FilmRankedReply, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	key := FormatKeyRanked(request.Category.String(), request.Offset, request.Limit)

	ch := utils.GetSingleFlight().DoChan(key, func() (interface{}, error) {
		var result pb.FilmRankedReply

		response, err := s.redisRepository.LoadFilmsCache(key)

		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(key)
		}()
		if err == nil {
			result.RankedFilm = make([]*pb.FilmRanked, len(response))
			for i, film := range response {
				result.RankedFilm[i] = &pb.FilmRanked{Id: film.Id, Title: film.Title, Popularity: film.Popularity}
			}
			return &result, nil
		} else {
			films, err := s.mongodbRepository.GetRankedFilms(ctx, request)
			if err != nil {
				return nil, err
			}

			result.RankedFilm = make([]*pb.FilmRanked, len(films))
			for i, film := range films {
				result.RankedFilm[i] = &pb.FilmRanked{Id: film.Id, Title: film.Title, Popularity: film.Popularity}
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
		return &pb.FilmRankedReply{}, ctx.Err()
	case ret := <-ch:
		return ret.Val.(*pb.FilmRankedReply), ret.Err
	}
}
