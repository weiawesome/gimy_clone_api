package routes

import (
	"api_affair/api/handler/ad"
	adService "api_affair/service/ad"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitAdRoutes(r *gin.RouterGroup, c *grpc.ClientConn) {
	adTypeRoute := ":" + utils.GetAdTypeRouteParameter()

	r.GET("/"+adTypeRoute, ad.Handler{Service: adService.NewAdService(c)}.GetAd)
}
