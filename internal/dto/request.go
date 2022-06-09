package dto

type (
	CreateWalletRequest struct {
		Balance     int    `json:"balance"`
		PhoneNumber string `json:"phone_number"`
	} //@name CreateWalletRequest

	GetBalanceRequest struct {
		ID string `json:"id"`
	} //@name GetBalanceRequest

	ChargeWalletRequest struct {
		ID     string `json:"id"`
		Amount int    `json:"amount"`
	} //@name ChargeWalletRequest

	TransactionRequest struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
	} //@name TransactionRequest
)
