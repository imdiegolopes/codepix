package usecase

import (
	model "codepix/domain/model"
	"fmt"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, fmt.Errorf("PixUseCase: No account was found for given account ID %s", accountId)
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, fmt.Errorf("PixUseCase: No pix key was created")
	}

	p.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("PixUseCase: Unable to create new key at the moment")
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, fmt.Errorf("PixUsedCase: No Pix Key was found for given Key %s and Kind %s", key, kind)
	}

	return pixKey, nil
}
