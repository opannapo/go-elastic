package services

type IndexService interface {
	GetInfo() (result interface{}, err error)
	GetCat() (result interface{}, err error)
}

type UserService interface {
	GetAll() (result interface{}, err error)
	GetById(id int64) (result interface{}, err error)
	CreateOne(data interface{}) (result interface{}, err error)
	UpdateOne(id int64, data interface{}) (result interface{}, err error)
	ReplaceOne(id int64, data interface{}) (result interface{}, err error)
}
