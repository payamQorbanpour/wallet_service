package usecase

import (
	"context"
	"errors"
	"sync"

	"github.com/go-kit/kit/log"
)

type repo struct {
	db     map[string]int
	logger log.Logger
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		db:     map[string]int{},
		logger: logger,
	}
}

func (repo *repo) CreateWallet(ctx context.Context, wallet Wallet) error {
	if repo.checkWalletExistance(ctx, wallet.ID) {
		return errors.New("Wallet with this phone number already exists")
	}

	repo.db[wallet.ID] = wallet.Balance

	return nil
}

func (repo *repo) GetWallet(ctx context.Context, id string) (Wallet, error) {
	var mutex sync.Mutex

	mutex.Lock()
	defer mutex.Unlock()

	wallet := Wallet{
		ID:      id,
		Balance: repo.db[id],
	}

	if !repo.checkWalletExistance(ctx, id) {
		wallet.Balance = -1
		return wallet, errors.New("Wallet with this phone number does not exist")
	}

	return wallet, nil
}

func (repo *repo) ChargeWallet(ctx context.Context, id string, amount int) (Wallet, error) {
	wallet, err := repo.GetWallet(ctx, id)
	if err != nil {
		return wallet, err
	}

	wallet.Balance += amount
	repo.db[id] = wallet.Balance

	return wallet, nil
}

func (repo *repo) checkWalletExistance(ctx context.Context, id string) bool {
	if _, exists := repo.db[id]; exists {
		return true
	}

	return false
}
