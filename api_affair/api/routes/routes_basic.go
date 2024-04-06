package routes

import (
	"api_affair/api/handler/film"
	filmService "api_affair/service/film"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitBasicRoutes(r *gin.RouterGroup, c *grpc.ClientConn) {
	r.GET("", film.Handler{Service: filmService.NewFilmService(c)}.GetBasicFilms)
}
