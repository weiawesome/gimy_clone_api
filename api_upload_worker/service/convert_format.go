package service

import (
	"api_upload_worker/utils"
)

func convertFormat(filePath string, fileName string, translateName string) error {
	stream, err := utils.GetVideoStream(filePath, fileName)
	if err != nil {
		return err
	}
	return utils.TranslateToHLS(filePath, fileName, translateName, stream)
}
