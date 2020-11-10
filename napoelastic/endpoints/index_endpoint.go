package endpoints

import (
	"github.com/gin-gonic/gin"
	"napoelastic/napoelastic/services"
)

type IndexEndpoint struct {
	Group    *gin.RouterGroup
	Services services.IndexService
}

func InitIndexEndpoint(group *gin.RouterGroup, services services.IndexService) {
	instance := &IndexEndpoint{
		Group:    group,
		Services: services,
	}

	group.GET("/", instance.getInfo)
	group.GET("/cat", instance.getCat)
}

func (instance *IndexEndpoint) getInfo(ctx *gin.Context) {
	result, err := instance.Services.GetInfo()
	if err != nil {
		OutFailed(ctx, 200, err.Error())
		return
	}

	OutOk(ctx, result)
}

func (instance *IndexEndpoint) getCat(ctx *gin.Context) {
	result, err := instance.Services.GetCat()
	if err != nil {
		OutFailed(ctx, 200, err.Error())
		return
	}

	OutOk(ctx, result)
}
