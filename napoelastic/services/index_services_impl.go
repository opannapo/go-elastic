package services

import "napoelastic/napoelastic/repository"

type IndexServicesImpl struct {
	Repository repository.IndexRepository
}

func NewIndexServicesImpl(repository repository.IndexRepository) IndexService {
	return &IndexServicesImpl{
		Repository: repository,
	}
}

func (i *IndexServicesImpl) GetCat() (result interface{}, err error) {
	result, err = i.Repository.GetCat()
	return
}

func (i *IndexServicesImpl) GetInfo() (result interface{}, err error) {
	result, err = i.Repository.GetInfo()
	return
}
