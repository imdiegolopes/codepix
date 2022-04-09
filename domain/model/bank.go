package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base     `valid:"required"`
	Name     string     `json:"name" gorm:"type:varchar(20)" valid:"notnull"`
	Code     string     `json:"code" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

// Isso é um método no Go.
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

// Isso é uma função no Go.
func NewBank(name string, code string) (*Bank, error) {
	bank := Bank{
		Name: name,
		Code: code,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	bank.UpdatedAt = time.Now()

	// Retorna o error na variável de `err`
	err := bank.isValid()
	// Se houver algum error na camada de Model, então retorna algo.
	if err != nil {
		return nil, err
	}

	// Por ser um ponteiro, é necessário retornar a referência em memória e não o objeto em si.
	return &bank, nil
}
