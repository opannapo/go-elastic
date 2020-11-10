package napoelastic

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	Config
	ConfigFile string
}

func (instance *App) Run() {
	viper.SetConfigFile(instance.ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	instance.Config = Config{
		Mode:    viper.GetString("mode"),
		Server:  &Server{Address: viper.GetString("server.address")},
		Context: &Context{Timeout: viper.GetInt("context.timeout")},
		Elastic: &Elastic{
			Address: viper.GetStringSlice("elastic.address"),
		},
	}
	instance.configPrintOut()

	db, err := instance.setupDb()
	if err != nil {
		panic(err)
		return
	}

	router, err := instance.setupRouter(db)
	if err != nil {
		panic(err)
		return
	}
	_ = router.Run(instance.Config.Server.Address)
	if err != nil {
		panic(err)
		return
	}
}

func (instance *App) setupDb() (db *DB, err error) {
	db = new(DB)
	err = db.setupElasticClient(&instance.Config)

	if err != nil {
		panic(err)
		return
	}

	esClient, err := db.createNewClient()
	if err != nil {
		panic(err)
		return
	}

	db.EsClient = esClient
	return
}

func (instance *App) setupRouter(db *DB) (routerEngine *gin.Engine, err error) {
	routerEngine = initRouter(&instance.Config, db)
	return
}
