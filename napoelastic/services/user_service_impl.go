package services

import (
	"napoelastic/napoelastic/repository"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{Repository: repository}
}

func (instance *UserServiceImpl) GetAll() (result interface{}, err error) {
	result, err = instance.Repository.GetAll()
	return
}
