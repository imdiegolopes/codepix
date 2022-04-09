package model_test

import (
	model "codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(name, code)

	ownerName := "Diego Lopes"
	accountNumber := "0001-abc"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil((account.ID)))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.OwnerName, ownerName)
	require.NotEmpty(t, account.CreatedAt)
	require.Equal(t, bank.Name, name)
	require.Equal(t, bank.Code, code)
	require.NotEmpty(t, bank.CreatedAt)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
