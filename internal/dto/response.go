package dto

import "time"

type (
	Wallet struct {
		ID        string `json:"id,omitempty"`
		Balance   int    `json:"balance"`
		CreatedAt time.Time
	}

	CreateWalletResponse struct {
		ID string `json:"id"`
	}

	GetBalanceResponse struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	}
)
