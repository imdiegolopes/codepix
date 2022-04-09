package model_test

import (
	model "codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPixKey(t *testing.T) {
	bankName := "NuBank"
	code := "291"
	bank, err := model.NewBank(bankName, code)

	accountNumber := "123456-abc"
	ownerName := "Diego Lopes"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	kind := "cpf"
	key := "000000000-00"

	pixkey, err := model.NewPixKey(kind, account, key)

	require.Nil(t, err)
	require.NotNil(t, bank)
	require.NotNil(t, account)
	require.NotNil(t, bank)
	require.NotNil(t, pixkey)
	require.NotEmpty(t, uuid.FromStringOrNil(pixkey.ID))
	require.Equal(t, pixkey.Key, key)
	require.Equal(t, pixkey.Kind, kind)
	require.Equal(t, pixkey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
