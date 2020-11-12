package es

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"napoelastic/napoelastic/repository"
)

type UserRepositoryImpl struct {
	ESClient *elasticsearch.Client
}

func NewUserRepositoryImpl(esClient *elasticsearch.Client) repository.UserRepository {
	return &UserRepositoryImpl{ESClient: esClient}
}

func (instance *UserRepositoryImpl) GetAll() (result interface{}, err error) {
	es := instance.ESClient

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("user"),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	err = instance.parsingHitsAsArray(res, &result)
	return
}
