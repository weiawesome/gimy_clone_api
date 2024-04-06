package film

import (
	"api_affair/api/response/failure"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetSearchCelebrityFilms(c *gin.Context) {
	content := c.DefaultQuery(utils.GetContentParameter(), utils.GetDefaultValue())
	offsetValue := c.DefaultQuery(utils.GetOffsetParameter(), utils.GetDefaultValue())
	limitValue := c.DefaultQuery(utils.GetLimitParameter(), utils.GetDefaultValue())

	offset := GetOffset(offsetValue)
	limit := GetLimit(limitValue)

	films, err := h.Service.GetSearchCelebrityFilms(content, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}
func (h Handler) GetSearchContentFilms(c *gin.Context) {
	content := c.DefaultQuery(utils.GetContentParameter(), utils.GetDefaultValue())
	offsetValue := c.DefaultQuery(utils.GetOffsetParameter(), utils.GetDefaultValue())
	limitValue := c.DefaultQuery(utils.GetLimitParameter(), utils.GetDefaultValue())

	offset := GetOffset(offsetValue)
	limit := GetLimit(limitValue)

	films, err := h.Service.GetSearchContentFilms(content, offset, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}
