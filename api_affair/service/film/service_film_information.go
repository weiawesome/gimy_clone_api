package film

import (
	"api_affair/api/response/film"
	pb "api_affair/proto/film_service"
	"api_affair/utils"
	"context"
)

func (s *filmService) GetFilmInformation(id string) (film.FilmInformation, error) {
	var result film.FilmInformation

	client := pb.NewFilmClient(s.c)
	request := pb.FilmSpecificRequest{Id: id}
	filmResult, err := client.GetSpecificFilm(context.Background(), &request)
	if err != nil {
		return result, err
	}

	go func() {
		_, err := client.AddFilmPopularity(context.Background(), &request)
		if err != nil {
			return
		}
	}()

	cdnAddress := utils.EnvCDNAddress()

	result = film.FilmInformation{
		Id:           filmResult.Id,
		Title:        filmResult.Title,
		Resource:     cdnAddress + filmResult.Resource,
		State:        filmResult.State,
		Type:         filmResult.Type.String(),
		Category:     filmResult.Category.String(),
		Actors:       filmResult.Actors,
		Directors:    filmResult.Directors,
		Location:     filmResult.Location.String(),
		ReleaseYear:  filmResult.ReleaseYear,
		UpdateTime:   filmResult.UpdateTime.AsTime(),
		Popularity:   filmResult.Popularity,
		Introduction: filmResult.Introduction,
	}
	if result.Actors == nil {
		result.Actors = []string{}
	}
	if result.Directors == nil {
		result.Directors = []string{}
	}
	return result, nil
}

func (s *filmService) GetFilmRouteInformation(id string) (film.FilmRouteInformation, error) {
	var result film.FilmRouteInformation

	client := pb.NewFilmClient(s.c)
	request := pb.FilmSpecificRequest{Id: id}
	routes, err := client.GetSpecificFilmRoutes(context.Background(), &request)
	if err != nil {
		return result, err
	}

	result.FilmRoutes = make([]film.FilmRoute, len(routes.Routes))
	for i, route := range routes.Routes {
		result.FilmRoutes[i] = film.FilmRoute{
			Route:    route.Route,
			Episodes: route.Episodes,
		}
		if result.FilmRoutes[i].Episodes == nil {
			result.FilmRoutes[i].Episodes = []string{}
		}
	}
	return result, nil
}
