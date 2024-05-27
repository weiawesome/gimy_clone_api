package route

import (
	"api_media/api/handler"
	"api_media/repository"
	"api_media/service"
	"api_media/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitResourceRoutes(route *gin.RouterGroup, r *repository.MinIORepository, c *grpc.ClientConn) {
	parameterFileBucket := ":" + utils.GetFileBucketParameter()
	parameterFile := ":" + utils.GetFileParameter()

	parameterPlayRoute := ":" + utils.GetPlayRouteParameter()
	parameterFileID := ":" + utils.GetFileIDParameter()
	parameterFileEpisode := ":" + utils.GetFileEpisodeParameter()
	parameterFileKey := ":" + utils.GetFileKeyParameter()

	route.GET("/media/"+parameterFileBucket+"/"+parameterFile, handler.MediaHandler{Service: service.NewMediaService(*r, c)}.GetMedia)
	route.GET("/media_list/"+parameterPlayRoute+"/"+parameterFileID+"/"+parameterFileEpisode, handler.MediaHandler{Service: service.NewMediaService(*r, c)}.GetPlayList)
	route.GET("/media_list/"+parameterPlayRoute+"/"+parameterFileID+"/"+parameterFileEpisode+"/"+parameterFileKey, handler.MediaHandler{Service: service.NewMediaService(*r, c)}.GetKey)
}
