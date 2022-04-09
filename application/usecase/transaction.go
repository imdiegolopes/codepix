package usecase

import (
	model "codepix/domain/model"
	"fmt"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixKeyRepository      model.PixKeyRepositoryInterface
}

func (t *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := t.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, fmt.Errorf("TransactionUseCase: Account was not found")
	}

	pixKey, err := t.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, fmt.Errorf("TransactionUseCase: Transaction was not found")
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description, "")

	if err != nil {
		return nil, fmt.Errorf("TransactionUseCase: Transaction couldn't be formatted")
	}

	t.TransactionRepository.Save(transaction)

	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, fmt.Errorf("TransactionUseCase: Unable to register transaction")
}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("TransactionUseCase: Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("TransactionUseCase: Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil

}
