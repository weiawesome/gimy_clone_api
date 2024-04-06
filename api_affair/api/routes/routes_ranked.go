package routes

import (
	"api_affair/api/handler/film"
	filmService "api_affair/service/film"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitRankedRoutes(r *gin.RouterGroup, c *grpc.ClientConn) {
	filmCategoryRoute := ":" + utils.GetFilmCategoryRouteParameter()

	r.GET(filmCategoryRoute, film.Handler{Service: filmService.NewFilmService(c)}.GetRankedFilms)
}
