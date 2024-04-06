package utils

import "os"

const (
	defaultMinIOHost            = "localhost"
	defaultMinIOPort            = "9000"
	defaultMinIOAccessKeyID     = "DefaultUser"
	defaultMinIOAccessKeySecret = "DefaultPassword"
	defaultMinIOToken           = ""

	defaultAdServiceHost = "localhost"
	defaultAdServicePort = "8083"

	defaultFilmServiceHost = "localhost"
	defaultFilmServicePort = "9090"

	defaultKafkaHost      = "localhost"
	defaultKafkaPort      = "9092"
	defaultKafkaFilmTopic = "Default-Film-Topic"
	defaultKafkaAdTopic   = "Default-Ad-Topic"
)

func EnvMinIOAddress() string {
	var host string
	var port string
	if host = os.Getenv("MINIO_HOST"); len(host) == 0 {
		host = defaultMinIOHost
	}
	if port = os.Getenv("MINIO_PORT"); len(port) == 0 {
		port = defaultMinIOPort
	}
	return host + ":" + port
}
func EnvMinIOAccessKeyID() string {
	var id string
	if id = os.Getenv("MINIO_ACCESS_KEY_ID"); len(id) == 0 {
		id = defaultMinIOAccessKeyID
	}
	return id
}
func EnvMinIOAccessKeySecret() string {
	var secret string
	if secret = os.Getenv("MINIO_ACCESS_KEY_SECRET"); len(secret) == 0 {
		secret = defaultMinIOAccessKeySecret
	}
	return secret
}
func EnvMinIOAccessKeyToken() string {
	var token string
	if token = os.Getenv("MINIO_TOKEN"); len(token) == 0 {
		token = defaultMinIOToken
	}
	return token
}

func EnvAdServiceAddress() string {
	var host string
	var port string
	if host = os.Getenv("AD_SERVICE_HOST"); len(host) == 0 {
		host = defaultAdServiceHost
	}
	if port = os.Getenv("AD_SERVICE_PORT"); len(port) == 0 {
		port = defaultAdServicePort
	}
	return host + ":" + port
}
func EnvFilmServiceAddress() string {
	var host string
	var port string
	if host = os.Getenv("FILM_SERVICE_HOST"); len(host) == 0 {
		host = defaultFilmServiceHost
	}
	if port = os.Getenv("FILM_SERVICE_PORT"); len(port) == 0 {
		port = defaultFilmServicePort
	}
	return host + ":" + port
}

func EnvKafkaAddress() string {
	var host string
	var port string
	if host = os.Getenv("KAFKA_HOST"); len(host) == 0 {
		host = defaultKafkaHost
	}
	if port = os.Getenv("KAFKA_PORT"); len(port) == 0 {
		port = defaultKafkaPort
	}
	return host + ":" + port
}
func EnvKafkaFilmTopic() string {
	var topic string
	if topic = os.Getenv("KAFKA_FILM_TOPIC"); len(topic) == 0 {
		topic = defaultKafkaFilmTopic
	}
	return topic
}
func EnvKafkaAdTopic() string {
	var topic string
	if topic = os.Getenv("KAFKA_AD_TOPIC"); len(topic) == 0 {
		topic = defaultKafkaAdTopic
	}
	return topic
}
