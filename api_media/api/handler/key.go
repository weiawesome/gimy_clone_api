package handler

import (
	"api_media/api/response/failure"
	"api_media/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h MediaHandler) GetKey(c *gin.Context) {
	fileRoute := c.Param(utils.GetPlayRouteParameter())
	fileKey := c.Param(utils.GetFileKeyParameter())
	media, err := h.Service.GetStreamMedia(fileRoute, fileKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
	}
	info, err := media.Stat()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
	}

	contentType := info.ContentType
	size := info.Size

	c.DataFromReader(http.StatusOK, size, contentType, media, nil)

}
