package handler

import (
	"api_upload/api/response/failure"
	responseMedia "api_upload/api/response/media"
	"api_upload/api/reuqest/media"
	"api_upload/service"
	"api_upload/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type FilmHandler struct {
	Service service.FilmService
}

func getFileExtension(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return ""
	}
	return strings.ToLower(filename[index:])
}

func (h FilmHandler) UploadFilm(c *gin.Context) {
	var film media.FilmInformation

	if err := c.ShouldBindJSON(&film); err != nil {
		c.JSON(http.StatusBadRequest, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	err = h.Service.CreateFilm(id, film)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, responseMedia.FilmAddInformation{Id: id.String()})
}
func (h FilmHandler) UploadFilmSearchEngine(c *gin.Context) {
	id := c.Param(utils.GetFileIDParameter())
	err := h.Service.UploadFilmSearchEngine(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}
func (h FilmHandler) UploadFilmImage(c *gin.Context) {
	id := c.Param(utils.GetFileIDParameter())
	file, handler, err := c.Request.FormFile(utils.GetFileParameter())
	if err != nil {
		c.JSON(http.StatusBadRequest, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	contentType := handler.Header.Get("Content-Type")
	size := handler.Size
	err = h.Service.UploadFilmImage(id, file, contentType, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}
func (h FilmHandler) UploadFilmResource(c *gin.Context) {
	route := c.Param(utils.GetPlayRouteParameter())
	id := c.Param(utils.GetFileIDParameter())
	episode := c.Param(utils.GetFileEpisodeParameter())
	state := c.DefaultQuery(utils.GetStateParameter(), utils.GetDefaultValue())
	file, handler, err := c.Request.FormFile(utils.GetFileParameter())
	if err != nil {
		c.JSON(http.StatusBadRequest, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	fileExtension := getFileExtension(handler.Filename)
	contentType := handler.Header.Get("Content-Type")
	size := handler.Size

	err = h.Service.UploadFilmResource(route, id, episode, file, fileExtension, contentType, size, state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failure.ServerError{Reason: err.Error()})
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}
