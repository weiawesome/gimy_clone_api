package ad

import (
	"api_affair/api/response/failure"
	"api_affair/proto/ad_service"
	"api_affair/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetAd(c *gin.Context) {
	adType := c.Param(utils.GetAdTypeRouteParameter())
	ad, err := h.Service.GetAd(ad_service.AdType(ad_service.AdType_value[adType]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, ad)
}
