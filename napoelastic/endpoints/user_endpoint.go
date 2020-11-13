package endpoints

import (
	"github.com/gin-gonic/gin"
	"napoelastic/napoelastic/services"
	"strconv"
)

type UserEndpoint struct {
	Group    *gin.RouterGroup
	Services *services.UserService
}

func InitUserEndpoint(group *gin.RouterGroup, services *services.UserService) {
	instance := &UserEndpoint{
		Group:    group,
		Services: services,
	}

	group.GET("/users", instance.getUsers)
	group.GET("/user/:id", instance.getUser)
}

func (instance *UserEndpoint) getUsers(ctx *gin.Context) {
	s := *instance.Services
	result, err := s.GetAll()
	if err != nil {
		OutFailed(ctx, 200, err.Error())
		return
	}

	OutOk(ctx, result)
}

func (instance *UserEndpoint) getUser(ctx *gin.Context) {
	s := *instance.Services

	id, _ := strconv.Atoi(ctx.Param("id"))
	result, err := s.GetById(int64(id))
	if err != nil {
		OutFailed(ctx, 200, err.Error())
		return
	}

	OutOk(ctx, result)
}
