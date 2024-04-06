package api

import (
	pb "api_ad_service/proto/ad_service"
	"api_ad_service/service"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer interface {
	GetAd(ctx context.Context, in *pb.GetAdvertisementRequest) (*pb.GetAdvertisementReply, error)
	SaveAd(ctx context.Context, in *pb.SaveAdvertisementRequest) (*emptypb.Empty, error)
}

type Server struct {
	pb.UnimplementedAdvertisementServer
}

func (s *Server) GetAd(ctx context.Context, in *pb.GetAdvertisementRequest) (*pb.GetAdvertisementReply, error) {
	result := pb.GetAdvertisementReply{}
	advertisement, err := service.GetService().Get(ctx, in)
	if err == nil {
		result.Bucket = advertisement.Bucket
		result.File = advertisement.File
	}
	return &result, nil
}
func (s *Server) SaveAd(ctx context.Context, in *pb.SaveAdvertisementRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.GetService().Save(ctx, in)
}
