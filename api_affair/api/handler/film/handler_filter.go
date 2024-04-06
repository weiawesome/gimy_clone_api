package film

import (
	"api_affair/api/response/failure"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetFilterFilms(c *gin.Context) {
	filmType := c.Param(utils.GetFilmTypeRouteParameter())
	category := c.DefaultQuery(utils.GetCategoryParameter(), utils.GetDefaultValue())
	location := c.DefaultQuery(utils.GetLocationParameter(), utils.GetDefaultValue())
	releaseYearValue := c.DefaultQuery(utils.GetReleaseYearParameter(), utils.GetDefaultValue())
	orderType := c.DefaultQuery(utils.GetOrderTypeParameter(), utils.GetDefaultValue())
	offsetValue := c.DefaultQuery(utils.GetOffsetParameter(), utils.GetDefaultValue())
	limitValue := c.DefaultQuery(utils.GetLimitParameter(), utils.GetDefaultValue())

	releaseYear := GetReleaseYear(releaseYearValue)
	offset := GetOffset(offsetValue)
	limit := GetLimit(limitValue)

	films, err := h.Service.GetFilterFilms(filmType, category, location, releaseYear, orderType, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}
