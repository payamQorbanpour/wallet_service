package dto

type (
	CreateWalletRequest struct {
		Balance     int    `json:"balance"`
		PhoneNumber string `json:"phone_number"`
	}

	GetBalanceRequest struct {
		ID string `json:"id"`
	}

	ChargeWalletRequest struct {
		ID     string `json:"id"`
		Amount int    `json:"amount"`
	}
)
