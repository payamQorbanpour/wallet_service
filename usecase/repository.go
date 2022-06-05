package usecase

import (
	"context"

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
	repo.db[wallet.ID] = wallet.Balance

	return nil
}

func (repo *repo) GetWallet(ctx context.Context, id string) (Wallet, error) {
	wallet := Wallet{
		ID:      id,
		Balance: repo.db[id],
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
