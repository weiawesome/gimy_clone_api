package routes

import (
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitRoutes(filmConn *grpc.ClientConn, adConn *grpc.ClientConn) *gin.Engine {

	route := gin.Default()

	router := route.Group("/api/" + utils.GetVersion() + "/affair")

	basicRouter := router.Group("")
	filmInformationRouter := router.Group("/film_information")
	filterRouter := router.Group("/filter")
	popularRouter := router.Group("/popular")
	rankedRouter := router.Group("/ranked")
	searchRouter := router.Group("/search")
	adRouter := router.Group("/ad")

	InitBasicRoutes(basicRouter, filmConn)
	InitFilmInformationRoutes(filmInformationRouter, filmConn)
	InitFilterRoutes(filterRouter, filmConn)
	InitPopularRoutes(popularRouter, filmConn)
	InitRankedRoutes(rankedRouter, filmConn)
	InitSearchRoutes(searchRouter, filmConn)
	InitAdRoutes(adRouter, adConn)

	return route
}
