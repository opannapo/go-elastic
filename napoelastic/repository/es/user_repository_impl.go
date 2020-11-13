package es

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"napoelastic/napoelastic/repository"
	"strconv"
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

func (instance *UserRepositoryImpl) GetById(id int64) (result interface{}, err error) {
	es := instance.ESClient
	res, err := es.GetSource("user", strconv.Itoa(int(id)))
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	err = instance.parsingSource(res, &result)
	return
}

func (instance *UserRepositoryImpl) CreateOne(data interface{}) (result interface{}, err error) {
	panic("implement me")
}

func (instance *UserRepositoryImpl) UpdateOne(id int64, data interface{}) (result interface{}, err error) {
	panic("implement me")
}

func (instance *UserRepositoryImpl) ReplaceOne(id int64, data interface{}) (result interface{}, err error) {
	panic("implement me")
}
