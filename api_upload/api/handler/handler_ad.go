package handler

import (
	"api_upload/api/response/failure"
	"api_upload/service"
	"api_upload/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AdHandler struct {
	Service service.AdService
}

func (h AdHandler) Upload(c *gin.Context) {
	adType := c.Request.FormValue(utils.GetAdTypeParameter())
	expiredTime := c.Request.FormValue(utils.GetExpiredTimeParameter())
	parsedTime, err := time.Parse(time.RFC3339, expiredTime)
	file, handler, err := c.Request.FormFile(utils.GetFileParameter())
	if err != nil {
		c.JSON(http.StatusBadRequest, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	fileExtension := getFileExtension(handler.Filename)
	contentType := handler.Header.Get("Content-Type")
	size := handler.Size
	err = h.Service.CreateAd(adType, file, fileExtension, contentType, size, parsedTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}
