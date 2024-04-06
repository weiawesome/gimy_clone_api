package routes

import (
	"api_affair/api/handler/film"
	filmService "api_affair/service/film"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitFilmInformationRoutes(r *gin.RouterGroup, c *grpc.ClientConn) {
	filmIdRoute := ":" + utils.GetFilmIdRouteParameter()
	r.GET("/"+filmIdRoute, film.Handler{Service: filmService.NewFilmService(c)}.GetFilmInformation)
	r.GET("/routes/"+filmIdRoute, film.Handler{Service: filmService.NewFilmService(c)}.GetFilmRouteInformation)
}
