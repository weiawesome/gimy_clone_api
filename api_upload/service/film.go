package service

import (
	"api_upload/api/reuqest/media"
	"api_upload/proto/film_service"
	"api_upload/repository/minio"
	"api_upload/utils"
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"mime/multipart"
)

type FilmService interface {
	CreateFilm(uuid.UUID, media.FilmInformation) error
	UploadFilmSearchEngine(string) error
	UploadFilmImage(string, multipart.File, string, int64) error
	UploadFilmResource(string, string, string, multipart.File, string, string, int64, string) error
}

type filmService struct {
	r minio.Repository
	c *grpc.ClientConn
	p message.Publisher
}

func NewFilmService(r minio.Repository, c *grpc.ClientConn, p message.Publisher) FilmService {
	return &filmService{r: r, c: c, p: p}
}

func (s *filmService) CreateFilm(id uuid.UUID, filmInformation media.FilmInformation) error {
	client := film_service.NewFilmClient(s.c)
	film := film_service.FilmSaveRequest{
		Id:           id.String(),
		Title:        filmInformation.Title,
		Resource:     "/api/v1/resource/media/" + utils.GetImageResourceBucket() + "/" + id.String(),
		State:        filmInformation.State,
		Type:         film_service.MediaType(film_service.MediaType_value[filmInformation.Type]),
		Category:     film_service.MediaCategory(film_service.MediaCategory_value[filmInformation.Category]),
		Actors:       filmInformation.Actors,
		Directors:    filmInformation.Directors,
		Location:     film_service.MediaLocation(film_service.MediaLocation_value[filmInformation.Location]),
		ReleaseYear:  filmInformation.ReleaseYear,
		Introduction: filmInformation.Introduction,
		Language:     filmInformation.Language,
	}
	_, err := client.SaveFilm(context.Background(), &film)
	return err
}
func (s *filmService) UploadFilmSearchEngine(id string) error {
	client := film_service.NewFilmClient(s.c)
	_, err := client.AddFilmToSearchEngine(context.Background(), &film_service.FilmSpecificRequest{Id: id})
	return err
}
func (s *filmService) UploadFilmImage(id string, file multipart.File, contentType string, size int64) error {
	bucket := utils.GetImageResourceBucket()
	return s.r.Create(file, bucket, id, size, contentType)
}
func (s *filmService) UploadFilmResource(route string, id string, episode string, file multipart.File, fileExtension string, contentType string, size int64, state string) error {
	description := utils.GetOriginalFileDescription()
	err := s.r.Create(file, route, id+"-"+episode+description+fileExtension, size, contentType)
	if err != nil {
		return err
	}
	film := utils.FilmInformation{Id: id, Route: route, Episode: episode, FileExtension: fileExtension, State: state}
	bytes, err := json.Marshal(film)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), bytes)
	return s.p.Publish(utils.EnvKafkaFilmTopic(), msg)
}
