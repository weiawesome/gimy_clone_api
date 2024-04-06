package film

import (
	"api_affair/api/response/failure"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetBasicFilms(c *gin.Context) {
	offsetValue := c.DefaultQuery(utils.GetOffsetParameter(), utils.GetDefaultValue())
	limitValue := c.DefaultQuery(utils.GetLimitParameter(), utils.GetDefaultValue())

	offset := GetOffset(offsetValue)
	limit := GetLimit(limitValue)

	films, err := h.Service.GetBasicFilms(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, films)
}
