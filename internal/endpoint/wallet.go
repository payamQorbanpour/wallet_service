package endpoint

import (
	"context"

	"wallet_service/internal/dto"
	"wallet_service/internal/pkg"

	"github.com/go-kit/kit/endpoint"
)

// @Summary Create wallet
// @ID create
// @Description Create wallet including wallet id and balance
// @Accept json
// @Produce json
// @Tags Create
// @Param request body dto.CreateWalletRequest true "Create request"
// @Success 200 {object} dto.CreateWalletResponse
// @Router /create [post]
// .
func makeCreateWalletEndpoint(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreateWalletRequest)
		id, err := s.CreateWallet(ctx, req.Balance, req.PhoneNumber)
		return dto.CreateWalletResponse{ID: id}, err
	}
}

// @Summary Get wallet
// @ID get
// @Description Get wallet data by wallet id
// @Accept json
// @Produce json
// @Tags Get
// @Param request body dto.GetBalanceRequest true "Get request"
// @Success 200 {object} dto.GetBalanceResponse
// @Failure 404 {object} dto.Error
// @Router /{id} [get]
// .
func makeGetWalletEndpoint(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.GetBalanceRequest)
		balance, err := s.GetBalance(ctx, req.ID)
		return dto.GetBalanceResponse{ID: req.ID, Balance: balance}, err
	}
}

// @Summary Charge wallet
// @ID charge
// @Description Charge wallet balance with given amount
// @Accept json
// @Produce json
// @Tags Charge
// @Param request body dto.ChargeWalletRequest true "Charge request"
// @Success 200 {object} dto.GetBalanceResponse
// @Failure 404 {object} dto.Error
// @Router /{id} [get]
// .
func makeChargeWalletEndpoint(s pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.ChargeWalletRequest)
		balance, err := s.ChargeWallet(ctx, req.ID, req.Amount)
		return dto.GetBalanceResponse{ID: req.ID, Balance: balance}, err
	}
}
