package film

import (
	"api_affair/api/response/failure"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetFilmInformation(c *gin.Context) {
	id := c.Param(utils.GetFilmIdRouteParameter())
	information, err := h.Service.GetFilmInformation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		return
	}
	c.JSON(http.StatusOK, information)
}
func (h Handler) GetFilmRouteInformation(c *gin.Context) {
	id := c.Param(utils.GetFilmIdRouteParameter())
	information, err := h.Service.GetFilmRouteInformation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		return
	}
	c.JSON(http.StatusOK, information)
}
