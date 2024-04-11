package routes

import (
	"api_upload/api/handler"
	"api_upload/repository/minio"
	"api_upload/service"
	"api_upload/utils"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitUploadMediaRoutes(route *gin.RouterGroup, r minio.Repository, c *grpc.ClientConn, p message.Publisher) {
	parameterPlayRoute := ":" + utils.GetPlayRouteParameter()
	parameterFileID := ":" + utils.GetFileIDParameter()
	parameterFileEpisode := ":" + utils.GetFileEpisodeParameter()

	route.POST("/",
		handler.FilmHandler{Service: service.NewFilmService(r, c, p)}.UploadFilm)
	route.DELETE("/"+parameterFileID,
		handler.FilmHandler{Service: service.NewFilmService(r, c, p)}.DeleteFilm)
	route.POST("/search_engine/"+parameterFileID,
		handler.FilmHandler{Service: service.NewFilmService(r, c, p)}.UploadFilmSearchEngine)
	route.POST("/image/"+parameterFileID,
		handler.FilmHandler{Service: service.NewFilmService(r, c, p)}.UploadFilmImage)
	route.POST("/resource/"+parameterPlayRoute+"/"+parameterFileID+"/"+parameterFileEpisode,
		handler.FilmHandler{Service: service.NewFilmService(r, c, p)}.UploadFilmResource)
	route.DELETE("/resource/"+parameterPlayRoute+"/"+parameterFileID+"/"+parameterFileEpisode,
		handler.FilmHandler{Service: service.NewFilmService(r, c, p)}.DeleteFilmResource)
}
