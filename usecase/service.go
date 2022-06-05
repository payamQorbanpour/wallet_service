package usecase

import "context"

type Service interface {
	CreateWallet(ctx context.Context, balance int, phoneNumber string) (id string, err error)
	GetBalance(ctx context.Context, id string) (balance int, err error)
	ChargeWallet(ctx context.Context, id string, amount int) (balance int, err error)
}
