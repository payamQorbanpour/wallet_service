package repository

import (
	"context"
	"sync"

	"wallet_service/internal/dto"
	"wallet_service/internal/model"
)

func (repo *Repo) CreateWallet(ctx context.Context, wallet dto.Wallet) error {
	if repo.checkWalletExistance(ctx, wallet.ID) {
		return model.ErrWalletExists
	}

	repo.DB[wallet.ID] = wallet.Balance

	return nil
}

func (repo *Repo) GetWallet(ctx context.Context, id string) (*dto.Wallet, error) {
	wallet := dto.Wallet{
		ID:      id,
		Balance: repo.DB[id],
	}

	if !repo.checkWalletExistance(ctx, id) {
		wallet.Balance = -1
		return nil, model.ErrWalletDoesNotExist
	}

	return &wallet, nil
}

func (repo *Repo) ChargeWallet(ctx context.Context, id string, amount int) (*dto.Wallet, error) {
	var mutex sync.Mutex

	mutex.Lock()
	defer mutex.Unlock()

	wallet, err := repo.GetWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	wallet.Balance += amount
	repo.DB[id] = wallet.Balance

	return wallet, nil
}

func (repo *Repo) Transaction(ctx context.Context, id, destinationID string, amount int) (*dto.Wallet, error) {
	var mutex sync.Mutex

	mutex.Lock()
	defer mutex.Unlock()

	wallet, err := repo.GetWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	destinationWallet, err := repo.GetWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	if !repo.checkWalletBalance(ctx, amount, *wallet) {
		return nil, model.ErrNotEnoghBalance
	}

	wallet.Balance -= amount
	repo.DB[id] = wallet.Balance

	destinationWallet.Balance += amount
	repo.DB[destinationID] = destinationWallet.Balance

	return destinationWallet, nil
}

func (repo *Repo) checkWalletExistance(ctx context.Context, id string) bool {
	if _, exists := repo.DB[id]; exists {
		return true
	}

	return false
}

func (repo *Repo) checkWalletBalance(ctx context.Context, amount int, wallet dto.Wallet) bool {
	return wallet.Balance > amount
}
