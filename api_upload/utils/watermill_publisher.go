package utils

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

var publisher message.Publisher

func InitPublisher() error {
	var err error

	address := EnvKafkaAddress()

	publisher, err = kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{address},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	return err
}
func GetPublisher() message.Publisher {
	return publisher
}
