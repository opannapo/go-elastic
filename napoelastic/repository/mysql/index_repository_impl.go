package mysql

import (
	"fmt"
	"napoelastic/napoelastic/repository"
)

type IndexRepositoryImpl struct {
	//todo Field for MysqlClient / Gorm
}

func NewIndexRepositoryImpl() repository.IndexRepository {
	return &IndexRepositoryImpl{}
}

func (instance *IndexRepositoryImpl) GetInfo() (result interface{}, err error) {
	err = fmt.Errorf("Mysql Repository Belum disetup")
	return
}

func (instance *IndexRepositoryImpl) GetCat() (result interface{}, err error) {
	err = fmt.Errorf("Mysql Repository Belum disetup")
	return
}
