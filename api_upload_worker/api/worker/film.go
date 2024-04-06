package worker

import (
	"api_upload_worker/api/request"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
)

func (w Worker) WorkFilm(message *message.Message) {
	var information request.FilmInformation
	var err error

	err = json.Unmarshal(message.Payload, &information)
	if err != nil {
		log.Println("error to unmarshal message with " + err.Error())
		return
	}
	err = w.Service.TranslateFilmResource(information)
	if err != nil {
		log.Println("error to translate film with " + err.Error())
		return
	}
	message.Ack()
}
