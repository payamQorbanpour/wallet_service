package dto

type (
	CreateWalletResponse struct {
		ID string `json:"id"`
	}

	GetBalanceResponse struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	}
)
