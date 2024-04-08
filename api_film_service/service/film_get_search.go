package service

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/model"
	"api_film_service/utils"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func (s filmService) GetSearchFilms(ctx context.Context, request *pb.FilmSearchRequest) (*pb.FilmSearchReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	var key string

	switch request.SearchType {
	case pb.SearchType_TITLE:
		key = FormatKeySearchContent(request.Content, request.Offset, request.Limit)
	case pb.SearchType_CELEBRITY:
		key = FormatKeySearchCelebrity(request.Content, request.Offset, request.Limit)
	}

	ch := utils.GetSingleFlight().DoChan(key, func() (interface{}, error) {
		var result pb.FilmSearchReply

		response, err := s.redisRepository.LoadFilmsCache(key)

		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(key)
		}()
		if err == nil {
			result.Results = make([]*pb.FilmSearchResult, len(response))
			for i, film := range response {
				result.Results[i] = &pb.FilmSearchResult{Id: film.Id, Title: film.Title, Resource: film.Resource, State: film.State, Category: pb.MediaCategory(pb.MediaCategory_value[film.Category]), Actors: film.Actors, Director: film.Directors, Location: pb.MediaLocation(pb.MediaLocation_value[film.Location]), Language: film.Language, UpdateTime: timestamppb.New(film.UpdateAt), Introduction: film.Introduction, ReleaseYear: film.ReleaseYear}
				if result.Results[i].Actors == nil {
					result.Results[i].Actors = []string{}
				}
				if result.Results[i].Director == nil {
					result.Results[i].Director = []string{}
				}
			}
			return &result, nil
		} else {
			var films []model.Film
			var err error
			switch request.SearchType {
			case pb.SearchType_TITLE:
				films, err = s.elasticsearchRepository.SearchFilmsGet(ctx, request)
				if err != nil {
					return &result, err
				}
			case pb.SearchType_CELEBRITY:
				films, err = s.mongodbRepository.GetSearchFilms(ctx, request)
				if err != nil {
					return &result, err
				}
			}
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
		return &pb.FilmSearchReply{}, ctx.Err()
	case ret := <-ch:
		return ret.Val.(*pb.FilmSearchReply), ret.Err
	}
}
