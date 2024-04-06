package routes

import (
	"api_upload/api/handler"
	"api_upload/repository/minio"
	"api_upload/service"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitUploadAdRoutes(route *gin.RouterGroup, r minio.Repository, c *grpc.ClientConn, p message.Publisher) {
	route.POST("",
		handler.AdHandler{Service: service.NewAdService(r, c, p)}.Upload,
	)
}
