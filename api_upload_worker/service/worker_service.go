package service

import (
	"api_upload_worker/api/request"
	"api_upload_worker/proto/ad_service"
	"api_upload_worker/proto/film_service"
	"api_upload_worker/repository"
	"api_upload_worker/utils"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkerService interface {
	TranslateAdResource(request.AdInformation) error
	TranslateFilmResource(request.FilmInformation) error
}
type workerService struct {
	r               repository.Repository
	adServiceConn   *grpc.ClientConn
	filmServiceConn *grpc.ClientConn
}

func NewWorkerService(r repository.Repository, adServiceConn *grpc.ClientConn, filmServiceCOnn *grpc.ClientConn) WorkerService {
	return workerService{r: r, adServiceConn: adServiceConn, filmServiceConn: filmServiceCOnn}
}

func (s workerService) TranslateFilmResource(information request.FilmInformation) error {
	description := utils.GetOriginalFileDescription()
	bucket := information.Route
	filePath := "./" + information.Route + "/" + information.Id + "/" + information.Episode
	fileName := information.Id + "-" + information.Episode + description + information.FileExtension
	translateName := information.Id + "-" + information.Episode
	err := downloadResource(s.r, bucket, filePath, fileName)
	if err != nil {
		return err
	}
	err = convertFormat(filePath, fileName, translateName)
	if err != nil {
		return err
	}
	err = uploadResource(s.r, bucket, filePath, fileName)
	if err != nil {
		return err
	}
	deleteResource(filePath)
	client := film_service.NewFilmClient(s.filmServiceConn)
	_, err = client.SaveFilmEpisode(context.Background(), &film_service.FilmSaveEpisodeRequest{Id: information.Id, Route: information.Route, Episode: information.Episode, State: information.State})
	if err != nil {
		return err
	}
	return nil
}
func (s workerService) TranslateAdResource(information request.AdInformation) error {
	description := utils.GetOriginalFileDescription()
	bucket := information.Bucket
	filePath := "./" + information.Bucket + "/" + information.Id
	fileName := information.Id + description + information.FileExtension
	translateName := information.Id
	err := downloadResource(s.r, bucket, filePath, fileName)
	if err != nil {
		return err
	}
	err = convertFormat(filePath, fileName, translateName)
	if err != nil {
		return err
	}
	err = uploadResource(s.r, bucket, filePath, fileName)
	if err != nil {
		return err
	}
	deleteResource(filePath)
	client := ad_service.NewAdvertisementClient(s.adServiceConn)
	_, err = client.SaveAd(context.Background(), &ad_service.SaveAdvertisementRequest{Type: ad_service.AdType_FILM, Bucket: bucket, File: translateName + ".m3u8", ExpireTime: timestamppb.New(information.ExpireTime)})
	if err != nil {
		return err
	}
	return nil
}
