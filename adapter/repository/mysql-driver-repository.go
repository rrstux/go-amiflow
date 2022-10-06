package repository

import "github.com/robertsan96/go-amiflow/contract"

type mysqlDriverRepository struct{}

func NewMysqlDriverRepository() contract.PersistentStorageContract {
	return &mysqlDriverRepository{}
}

func (repo *mysqlDriverRepository) CreatePersistentStorageGroup(name string) {}
func (repo *mysqlDriverRepository) DeletePersistentStorageGroup(name string) {}
