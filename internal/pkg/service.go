package pkg

import (
	"context"
	"time"

	"wallet_service/internal/dto"
	"wallet_service/internal/model"
	"wallet_service/internal/repository"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type walletService struct {
	repository repository.Repository
	logger     log.Logger
}

type Service interface {
	CreateWallet(ctx context.Context, balance int, phoneNumber string) (id string, err error)
	GetBalance(ctx context.Context, id string) (balance int, err error)
	ChargeWallet(ctx context.Context, id string, amount int) (balance int, err error)
	Transaction(ctx context.Context, id, destinationID string, amount int) (balance int, err error)
}

func NewService(repo repository.Repository, logger log.Logger) Service {
	return &walletService{
		repository: repo,
		logger:     logger,
	}
}

func (s *walletService) CreateWallet(ctx context.Context, balance int, phoneNumber string) (id string, err error) {
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

func (s *walletService) GetBalance(ctx context.Context, id string) (balance int, err error) {
	logger := log.With(s.logger, "method", "GetBalance")

	wallet, err := s.repository.GetWallet(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return -1, err
	}

	logger.Log("Get balance", id)

	return wallet.Balance, nil
}

func (s *walletService) ChargeWallet(ctx context.Context, id string, amount int) (balance int, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	wallet, err := s.repository.ChargeWallet(ctx, id, amount)
	if err != nil {
		level.Error(logger).Log("err", err)
		return -1, err
	}

	logger.Log("Charge wallet", wallet.Balance)

	return wallet.Balance, nil
}

func (s *walletService) Transaction(ctx context.Context, id, destinationID string, amount int) (balance int, err error) {
	logger := log.With(s.logger, "method", "ChargeWallet")

	if id == destinationID {
		return -1, model.ErrSelfTransaction
	}

	wallet, err := s.repository.Transaction(ctx, id, destinationID, amount)
	if err != nil {
		level.Error(logger).Log("err", err)
		return -1, err
	}

	logger.Log("Charge wallet", wallet.Balance)

	return wallet.Balance, nil
}
