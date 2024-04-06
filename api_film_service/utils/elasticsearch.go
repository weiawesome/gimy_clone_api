package utils

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var elasticClient *elasticsearch.Client

func InitElasticsearch() error {
	var err error
	config := elasticsearch.Config{
		Addresses: []string{
			EnvElasticsearchAddress(),
		},
		Username: EnvElasticsearchUser(),
		Password: EnvElasticsearchPassword(),
	}
	elasticClient, err = elasticsearch.NewClient(config)
	return err
}
func GetElasticsearchClient() *elasticsearch.Client {
	return elasticClient
}
