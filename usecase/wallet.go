package usecase

import (
	"context"
	"time"
)

type Wallet struct {
	ID        string `json:"id,omitempty"`
	Balance   int    `json:"balance"`
	CreatedAt time.Time
}

type Repository interface {
	CreateWallet(ctx context.Context, wallet Wallet) error
	GetWallet(ctx context.Context, id string) (Wallet, error)
	ChargeWallet(ctx context.Context, id string, amount int) (Wallet, error)
}
