package service

import (
	pb "api_media/proto/ad_service"
	"context"
)

func (s mediaService) GetAd() (string, string, error) {
	client := pb.NewAdvertisementClient(s.c)
	ctx := context.Background()
	ad, err := client.GetAd(ctx, &pb.GetAdvertisementRequest{Type: pb.AdType_FILM})
	return ad.Bucket, ad.File, err
}
