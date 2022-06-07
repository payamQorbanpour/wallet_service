package repository

import (
	"context"
	"wallet_service/internal/dto"

	"github.com/go-kit/kit/log"
)

type Repo struct {
	DB     map[string]int
	logger log.Logger
}

type Repository interface {
	CreateWallet(ctx context.Context, wallet dto.Wallet) error
	GetWallet(ctx context.Context, id string) (dto.Wallet, error)
	ChargeWallet(ctx context.Context, id string, amount int) (dto.Wallet, error)
}

func NewRepo(logger log.Logger) Repository {
	return &Repo{
		DB:     map[string]int{},
		logger: logger,
	}
}

func (repo *Repo) checkWalletExistance(ctx context.Context, id string) bool {
	if _, exists := repo.DB[id]; exists {
		return true
	}

	return false
}
