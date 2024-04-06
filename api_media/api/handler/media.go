package handler

import (
	"api_media/api/response/failure"
	"api_media/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h MediaHandler) GetMedia(c *gin.Context) {
	fileBucket := c.Param(utils.GetFileBucketParameter())
	file := c.Param(utils.GetFileParameter())
	media, err := h.Service.GetStreamMedia(fileBucket, file)
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
