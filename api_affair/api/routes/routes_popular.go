package routes

import (
	"api_affair/api/handler/film"
	filmService "api_affair/service/film"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitPopularRoutes(r *gin.RouterGroup, c *grpc.ClientConn) {
	filmTypeRoute := ":" + utils.GetFilmTypeRouteParameter()
	filmCategoryRoute := ":" + utils.GetFilmCategoryRouteParameter()

	r.GET("/type/"+filmTypeRoute, film.Handler{Service: filmService.NewFilmService(c)}.GetPopularTypeFilms)
	r.GET("/category/"+filmCategoryRoute, film.Handler{Service: filmService.NewFilmService(c)}.GetPopularCategoryFilms)
}
