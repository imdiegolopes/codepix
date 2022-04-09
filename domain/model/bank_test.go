package model_test

import (
	model "codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBank(t *testing.T) {
	name := "Banco Itaú"
	code := "341"

	bank, err := model.NewBank(name, code)

	require.Nil(t, err)
	require.NotNil(t, bank)
	require.NotEmpty(t, uuid.FromStringOrNil(bank.ID))
	require.Equal(t, bank.Name, name)
	require.Equal(t, bank.Code, code)
	require.NotNil(t, bank.CreatedAt)
	require.NotNil(t, bank.UpdatedAt)

	bank1, err := model.NewBank("", "")

	require.NotNil(t, err)
	require.Nil(t, bank1)
}
