package usecase

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

func (s service) CreateWallet(ctx context.Context, balance int, phoneNumber string) (id string, err error) {
	logger := log.With(s.logger, "method", "CreateWallet")

	wallet := Wallet{
		ID:        phoneNumber,
		Balance:   balance,
		CreatedAt: time.Time(time.Now()),
	}

	if err := s.repository.CreateWallet(ctx, wallet); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create wallet", phoneNumber)

	return phoneNumber, nil
}

func (s service) GetBalance(ctx context.Context, id string) (balance int, err error) {
	logger := log.With(s.logger, "method", "GetBalance")

	wallet, err := s.repository.GetWallet(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return -1, err
	}

	logger.Log("Get balance", id)

	return wallet.Balance, nil
}

func (s service) ChargeWallet(ctx context.Context, id string, amount int) (balance int, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	wallet, err := s.repository.ChargeWallet(ctx, id, amount)
	if err != nil {
		level.Error(logger).Log("err", err)
		return -1, err
	}

	logger.Log("Charge wallet", wallet.Balance)

	return wallet.Balance, nil
}
