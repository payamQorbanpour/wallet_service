package endpoint

import (
	"wallet_service/internal/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateWallet endpoint.Endpoint
	GetWallet    endpoint.Endpoint
	ChargeWallet endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateWallet: makeCreateWalletEndpoint(s),
		GetWallet:    makeGetWalletEndpoint(s),
		ChargeWallet: makeChargeWalletEndpoint(s),
	}
}
