package worker

import (
	"api_upload_worker/api/request"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
)

func (w Worker) WorkAd(message *message.Message) {
	var information request.AdInformation
	var err error

	err = json.Unmarshal(message.Payload, &information)
	if err != nil {
		log.Println("error to unmarshal message with " + err.Error())
		return
	}
	err = w.Service.TranslateAdResource(information)
	if err != nil {
		log.Println("error to translate ad with " + err.Error())
		return
	}
	message.Ack()
}
