package es

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

func (instance *UserRepositoryImpl) parsingHitsAsArray(res *esapi.Response, result *interface{}) (err error) {
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	var js map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&js); err != nil {
		log.Fatalf("Error parsing response to js: %s", err)
	}

	*result = js["hits"].(map[string]interface{})["hits"]
	return
}
