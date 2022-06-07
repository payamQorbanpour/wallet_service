package repository

import (
	"context"
	"errors"
	"sync"

	"wallet_service/internal/dto"
)

func (repo *Repo) CreateWallet(ctx context.Context, wallet dto.Wallet) error {
	if repo.checkWalletExistance(ctx, wallet.ID) {
		return errors.New("Wallet with this phone number already exists")
	}

	repo.DB[wallet.ID] = wallet.Balance

	return nil
}

func (repo *Repo) GetWallet(ctx context.Context, id string) (dto.Wallet, error) {
	var mutex sync.Mutex

	mutex.Lock()
	defer mutex.Unlock()

	wallet := dto.Wallet{
		ID:      id,
		Balance: repo.DB[id],
	}

	if !repo.checkWalletExistance(ctx, id) {
		wallet.Balance = -1
		return wallet, errors.New("Wallet with this phone number does not exist")
	}

	return wallet, nil
}

func (repo *Repo) ChargeWallet(ctx context.Context, id string, amount int) (dto.Wallet, error) {
	wallet, err := repo.GetWallet(ctx, id)
	if err != nil {
		return wallet, err
	}

	wallet.Balance += amount
	repo.DB[id] = wallet.Balance

	return wallet, nil
}

func (repo *Repo) checkWalletExistance(ctx context.Context, id string) bool {
	if _, exists := repo.DB[id]; exists {
		return true
	}

	return false
}
