package es

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

func parsingSourceArray(res *esapi.Response, out interface{}) (err error) {
	if res.IsError() {
		err = fmt.Errorf(parseErrorResponse(res))
		return
	}

	var resJson map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&resJson); err != nil {
		log.Printf("Error parsing response to js: %s", err)
		return
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
		err = fmt.Errorf(parseErrorResponse(res))
		log.Printf("parsingSource err %#v", err)
		return
	}

	if err = json.NewDecoder(res.Body).Decode(&out); err != nil {
		log.Printf("Error parsing response to result: %s", err)
	}

	return
}

func parseErrorResponse(res *esapi.Response) string {
	var outErr map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&outErr); err != nil {
		log.Printf("Error parseErrorResponse : %s", err)
	}

	return outErr["error"].(map[string]interface{})["type"].(string)
}
