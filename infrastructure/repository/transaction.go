package repository

import (
	model "codepix/domain/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDB struct {
	DB *gorm.DB
}

func (r TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := r.DB.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	err := r.DB.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.DB.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("TransactionRepositoryDB: No transaction was found for given ID %s", id)
	}

	return &transaction, nil
}
