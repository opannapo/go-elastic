package repository

type IndexRepository interface {
	GetInfo() (result interface{}, err error)
	GetCat() (result interface{}, err error)
}
