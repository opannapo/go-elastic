package napoelastic

import (
	"fmt"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"os"
)

const msgEs7 = `- ElasticSearch-7
	└ Version			» %s  
`

type DB struct {
	Config   *Config
	EsClient *es7.Client
	EsConfig *es7.Config
}

func (db *DB) setupElasticClient(config *Config) (err error) {
	fmt.Println("\n\n---->	setupElasticConfig...")

	clusterAddr := config.Elastic.Address
	cfg := es7.Config{
		Addresses: clusterAddr,
		Username:  "",
		Password:  "",
		Logger: &estransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	}

	db.EsConfig = &cfg
	return
}

func (db *DB) createNewClient() (esClient *es7.Client, err error) {
	fmt.Println("\n\n---->	createElasticClient...")

	esClient, err = es7.NewClient(*db.EsConfig)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf(msgEs7, es7.Version)
		fmt.Println(esClient.Info())
	}

	return
}
