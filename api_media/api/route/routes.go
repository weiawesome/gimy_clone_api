package route

import (
	"api_media/repository"
	"api_media/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitRoutes(r *repository.MinIORepository, c *grpc.ClientConn) *gin.Engine {

	route := gin.Default()

	basicRouter := route.Group("/api/" + utils.GetVersion())

	ResourceRouter := basicRouter.Group("/resource")

	InitResourceRoutes(ResourceRouter, r, c)

	return route
}
