package endpoint

import (
	"context"

	"wallet_service/internal/dto"
	"wallet_service/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateWalletEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreateWalletRequest)
		id, err := s.CreateWallet(ctx, req.Balance, req.PhoneNumber)
		return dto.CreateWalletResponse{ID: id}, err
	}
}

func makeGetWalletEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetBalanceRequest)
		balance, err := s.GetBalance(ctx, req.ID)
		return dto.GetBalanceResponse{ID: req.ID, Balance: balance}, err
	}
}

func makeChargeWalletEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.ChargeWalletRequest)
		balance, err := s.ChargeWallet(ctx, req.ID, req.Amount)
		return dto.GetBalanceResponse{ID: req.ID, Balance: balance}, err
	}
}
