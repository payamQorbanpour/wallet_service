package dto

import "time"

type (
	General struct {
		Message   string `json:"message"`
		ErrorCode int    `json:"error_code"`
	}

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
