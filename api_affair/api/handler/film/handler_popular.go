package film

import (
	"api_affair/api/response/failure"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetPopularTypeFilms(c *gin.Context) {
	filmType := c.Param(utils.GetFilmTypeRouteParameter())
	offsetValue := c.DefaultQuery(utils.GetOffsetParameter(), utils.GetDefaultValue())
	limitValue := c.DefaultQuery(utils.GetLimitParameter(), utils.GetDefaultValue())

	offset := GetOffset(offsetValue)
	limit := GetLimit(limitValue)

	films, err := h.Service.GetPopularTypeFilms(filmType, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		return
	}
	c.JSON(http.StatusOK, films)

}
func (h Handler) GetPopularCategoryFilms(c *gin.Context) {
	filmCategory := c.Param(utils.GetFilmCategoryRouteParameter())
	offsetValue := c.DefaultQuery(utils.GetOffsetParameter(), utils.GetDefaultValue())
	limitValue := c.DefaultQuery(utils.GetLimitParameter(), utils.GetDefaultValue())

	offset := GetOffset(offsetValue)
	limit := GetLimit(limitValue)

	films, err := h.Service.GetPopularCategoryFilms(filmCategory, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}
