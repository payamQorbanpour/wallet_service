package endpoint

import (
	"wallet_service/internal/pkg"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateWallet endpoint.Endpoint
	GetWallet    endpoint.Endpoint
	ChargeWallet endpoint.Endpoint
	Transaction  endpoint.Endpoint
}

func MakeEndpoints(s pkg.Service) Endpoints {
	return Endpoints{
		CreateWallet: makeCreateWalletEndpoint(s),
		GetWallet:    makeGetWalletEndpoint(s),
		ChargeWallet: makeChargeWalletEndpoint(s),
		Transaction:  makeTransactionEndpoint(s),
	}
}
