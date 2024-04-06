package routes

import (
	"api_upload/repository/minio"
	"api_upload/utils"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitRoutes(r minio.Repository, filmConn *grpc.ClientConn, adConn *grpc.ClientConn, publisher message.Publisher) *gin.Engine {

	route := gin.Default()

	basicRouter := route.Group("/api/" + utils.GetVersion() + "/upload")

	uploadMediaRouter := basicRouter.Group("/media")
	uploadAdRouter := basicRouter.Group("/ad")

	InitUploadMediaRoutes(uploadMediaRouter, r, filmConn, publisher)
	InitUploadAdRoutes(uploadAdRouter, r, adConn, publisher)

	return route
}
