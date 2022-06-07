package pkg

import (
	"context"
	"time"

	"wallet_service/internal/dto"
	"wallet_service/internal/repository"
	"wallet_service/internal/service"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type walletService struct {
	repository repository.Repository
	logger     log.Logger
}

func NewService(repo repository.Repository, logger log.Logger) service.Service {
	return &walletService{
		repository: repo,
		logger:     logger,
	}
}

func (s walletService) CreateWallet(ctx context.Context, balance int, phoneNumber string) (id string, err error) {
	logger := log.With(s.logger, "method", "CreateWallet")

	wallet := dto.Wallet{
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

func (s walletService) GetBalance(ctx context.Context, id string) (balance int, err error) {
	logger := log.With(s.logger, "method", "GetBalance")

	wallet, err := s.repository.GetWallet(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return -1, err
	}

	logger.Log("Get balance", id)

	return wallet.Balance, nil
}

func (s walletService) ChargeWallet(ctx context.Context, id string, amount int) (balance int, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	wallet, err := s.repository.ChargeWallet(ctx, id, amount)
	if err != nil {
		level.Error(logger).Log("err", err)
		return -1, err
	}

	logger.Log("Charge wallet", wallet.Balance)

	return wallet.Balance, nil
}
