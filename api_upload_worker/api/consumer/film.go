package consumer

import (
	"api_upload_worker/api/worker"
	"api_upload_worker/repository"
	"api_upload_worker/service"
	"api_upload_worker/utils"
	"context"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"strconv"
)

func filmProcess(messages <-chan *message.Message) {
	for msg := range messages {
		go worker.Worker{Service: service.NewWorkerService(repository.NewRepository(), utils.GetAdServiceConnection(), utils.GetFilmServiceConnection())}.WorkFilm(msg)
	}
}

func InitFilmConsumers() error {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	address := utils.EnvKafkaAddress()
	consumerGroup := utils.EnvKafkaConsumerGroup()
	topic := utils.EnvKafkaFilmTopic()

	replicasVal := utils.EnvFilmConsumerReplicas()

	replicas, err := strconv.Atoi(replicasVal)
	if err != nil {
		replicas = utils.GetDefaultReplicas()
	}

	for i := 0; i < replicas; i++ {
		subscriber, err := kafka.NewSubscriber(
			kafka.SubscriberConfig{
				Brokers:               []string{address},
				Unmarshaler:           kafka.DefaultMarshaler{},
				OverwriteSaramaConfig: saramaSubscriberConfig,
				ConsumerGroup:         consumerGroup,
			},
			watermill.NewStdLogger(false, false),
		)
		if err != nil {
			return err
		}
		messages, err := subscriber.Subscribe(context.Background(), topic)
		if err != nil {
			return err
		}
		go filmProcess(messages)
	}
	return nil
}
