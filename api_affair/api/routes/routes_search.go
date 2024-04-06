package routes

import (
	"api_affair/api/handler/film"
	filmService "api_affair/service/film"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitSearchRoutes(r *gin.RouterGroup, c *grpc.ClientConn) {
	r.GET("/content", film.Handler{Service: filmService.NewFilmService(c)}.GetSearchContentFilms)
	r.GET("/celebrity", film.Handler{Service: filmService.NewFilmService(c)}.GetSearchCelebrityFilms)
}
