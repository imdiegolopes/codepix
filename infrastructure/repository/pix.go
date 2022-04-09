package repository

import (
	model "codepix/domain/model"
	"fmt"

	"gorm.io/gorm"
)

// It's important to state the reposittories should be group by aggregated, not by entity
type PixKeyRespositoryDB struct {
	DB *gorm.DB
}

func (r PixKeyRespositoryDB) AddBank(bank *model.Bank) error {
	err := r.DB.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRespositoryDB) AddAccount(account *model.Account) error {
	err := r.DB.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRespositoryDB) RegisterKey(pixKey *model.PixKey) error {
	err := r.DB.Create(pixKey).Error

	if err != nil {
		return err
	}

	return err
}

func (r PixKeyRespositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.DB.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("PixKeyRespositoryDB: No key was found")
	}

	return &pixKey, nil
}

func (r PixKeyRespositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	r.DB.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("PixKeyRespositoryDB: No account was found")
	}

	return &account, nil
}

func (r PixKeyRespositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.DB.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("PixKeyRespositoryDB: No bank was found for given ID %s", id)
	}

	return &bank, nil
}
