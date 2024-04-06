package routes

import (
	"api_affair/api/handler/film"
	filmService "api_affair/service/film"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitFilterRoutes(r *gin.RouterGroup, c *grpc.ClientConn) {
	filmTypeRoute := ":" + utils.GetFilmTypeRouteParameter()

	r.GET("/"+filmTypeRoute, film.Handler{Service: filmService.NewFilmService(c)}.GetFilterFilms)
}
