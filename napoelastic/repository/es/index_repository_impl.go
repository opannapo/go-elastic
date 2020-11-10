package es

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"napoelastic/napoelastic/repository"
	"strings"
)

type IndexRepositoryImpl struct {
	ESClient *elasticsearch.Client
}

func NewIndexRepositoryImpl(eSClient *elasticsearch.Client) repository.IndexRepository {
	return &IndexRepositoryImpl{
		ESClient: eSClient,
	}
}

func (instance *IndexRepositoryImpl) GetCat() (result interface{}, err error) {
	res, err := instance.ESClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	// Deserialize the response into a map / interface.
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Println(strings.Repeat("~", 37))

	return
}

func (instance *IndexRepositoryImpl) GetInfo() (result interface{}, err error) {
	res, err := instance.ESClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	// Deserialize the response into a map / interface.
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Println(strings.Repeat("~", 37))

	return
}
