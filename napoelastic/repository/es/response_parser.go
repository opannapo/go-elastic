package es

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

func parsingSourceArray(res *esapi.Response, out interface{}) (err error) {
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	var resJson map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&resJson); err != nil {
		log.Fatalf("Error parsing response to js: %s", err)
	}

	var result []interface{}
	for _, hit := range resJson["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]

		jsonString, _ := json.Marshal(source)
		var tmp map[string]interface{}
		_ = json.Unmarshal(jsonString, &tmp)

		result = append(result, tmp)
		log.Printf("result = append %#v", result)
	}

	dbByte, _ := json.Marshal(result)
	_ = json.Unmarshal(dbByte, &out)

	return
}

func parsingSource(res *esapi.Response, out interface{}) (err error) {
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	if err = json.NewDecoder(res.Body).Decode(&out); err != nil {
		log.Fatalf("Error parsing response to result: %s", err)
	}

	log.Printf("out %#v", &out)

	return
}
