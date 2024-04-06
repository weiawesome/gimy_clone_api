package handler

import (
	"api_media/api/response/failure"
	"api_media/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h MediaHandler) GetPlayList(c *gin.Context) {
	playRoute := c.Param(utils.GetPlayRouteParameter())
	fileID := c.Param(utils.GetFileIDParameter())
	fileEpisode := c.Param(utils.GetFileEpisodeParameter())

	mediaList, err := h.Service.GetStreamMediaList(playRoute, fileID+"-"+fileEpisode+".m3u8")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, failure.ErrorResponse{Reason: err.Error()})
	}
	h.Service.RefreshBucketUrl(mediaList, playRoute)
	h.Service.RefreshMediaUrl(mediaList)

	bucket, file, err := h.Service.GetAd()
	if err == nil {
		mediaAdList, err := h.Service.GetStreamMediaList(bucket, file)
		if err == nil {
			h.Service.RefreshBucketUrl(mediaAdList, bucket)
			h.Service.RefreshMediaUrl(mediaAdList)
			h.Service.InsertAd(mediaList, mediaAdList)
		}
	}

	c.Data(http.StatusOK, "application/vnd.apple.mpegurl", []byte(mediaList.String()))
}
