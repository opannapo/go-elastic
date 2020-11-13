package services

import (
	"napoelastic/napoelastic/entities"
	"napoelastic/napoelastic/repository"
)

type UserServiceImpl struct {
	Repository *repository.UserRepository
}

func NewUserServiceImpl(repository *repository.UserRepository) UserService {
	return &UserServiceImpl{Repository: repository}
}

func (instance *UserServiceImpl) GetAll() (result []entities.UserEntity, err error) {
	r := *instance.Repository
	result, err = r.GetAll()
	return
}

func (instance *UserServiceImpl) GetById(id int64) (result *entities.UserEntity, err error) {
	r := *instance.Repository
	result, err = r.GetById(id)
	return
}

func (instance *UserServiceImpl) CreateOne(data interface{}) (result interface{}, err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) UpdateOne(id int64, data interface{}) (result interface{}, err error) {
	panic("implement me")
}

func (instance *UserServiceImpl) ReplaceOne(id int64, data interface{}) (result interface{}, err error) {
	panic("implement me")
}
