package usecase

import (
	"context"

	"wallet_service/usecase/dto"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateWallet endpoint.Endpoint
	GetWallet    endpoint.Endpoint
	ChargeWallet endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateWallet: makeCreateWalletEndpoint(s),
		GetWallet:    makeGetWalletEndpoint(s),
		ChargeWallet: makeChargeWalletEndpoint(s),
	}
}

func makeCreateWalletEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreateWalletRequest)
		id, err := s.CreateWallet(ctx, req.Balance, req.PhoneNumber)
		return dto.CreateWalletResponse{ID: id}, err
	}
}

func makeGetWalletEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetBalanceRequest)
		balance, err := s.GetBalance(ctx, req.ID)
		return dto.GetBalanceResponse{ID: req.ID, Balance: balance}, err
	}
}

func makeChargeWalletEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.ChargeWalletRequest)
		balance, err := s.ChargeWallet(ctx, req.ID, req.Amount)
		return dto.GetBalanceResponse{ID: req.ID, Balance: balance}, err
	}
}
