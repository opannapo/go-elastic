package napoelastic

import (
	"github.com/gin-gonic/gin"
	"napoelastic/napoelastic/endpoints"
	"napoelastic/napoelastic/repository/es"
	"napoelastic/napoelastic/services"
)

func initRouter(config *Config, db *DB) (router *gin.Engine) {
	gin.SetMode(config.Mode)

	router = gin.Default()
	router.NoRoute(func(context *gin.Context) {
		err := endpoints.ResError{
			Message: "Service Not Found",
			Code:    404,
		}
		res := endpoints.ResDefault{
			Error:   err,
			Data:    nil,
			Success: false,
		}
		context.JSON(200, res)
	})

	//initial dependency
	indexRepository := es.NewIndexRepositoryImpl(db.EsClient)
	indexServiceImpl := services.NewIndexServicesImpl(indexRepository)

	v1 := router.Group("api/v1")
	{
		endpoints.InitIndexEndpoint(v1, indexServiceImpl)
	}

	return
}
