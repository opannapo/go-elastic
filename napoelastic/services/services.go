package services

type IndexService interface {
	GetInfo() (result interface{}, err error)
	GetCat() (result interface{}, err error)
}
