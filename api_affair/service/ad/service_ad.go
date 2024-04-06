package ad

import (
	"api_affair/api/response/ad"
	"api_affair/proto/ad_service"
	"api_affair/utils"
	"context"
)

func (s *adService) GetAd(adType ad_service.AdType) (ad.AdInformation, error) {
	var result ad.AdInformation

	client := ad_service.NewAdvertisementClient(s.c)
	request := ad_service.GetAdvertisementRequest{Type: adType}
	advertisement, err := client.GetAd(context.Background(), &request)
	if err != nil {
		return result, err
	}

	cdnAddress := utils.EnvCDNAddress()

	result = ad.AdInformation{
		URL: cdnAddress + "/api/v1/resource/media/" + advertisement.Bucket + "/" + advertisement.File,
	}

	return result, nil
}
