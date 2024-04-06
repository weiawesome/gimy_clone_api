package ad

import (
	"api_affair/api/response/ad"
	"api_affair/proto/ad_service"
	"google.golang.org/grpc"
)

type Service interface {
	GetAd(ad_service.AdType) (ad.AdInformation, error)
}

type adService struct {
	c *grpc.ClientConn
}

func NewAdService(c *grpc.ClientConn) Service {
	return &adService{c: c}
}
