package dto

import "time"

type (
	Error struct {
		Message   string `json:"message"`
		ErrorCode int    `json:"error_code"`
	} //@name Error

	Wallet struct {
		ID        string `json:"id,omitempty"`
		Balance   int    `json:"balance"`
		CreatedAt time.Time
	} //@name Wallet

	CreateWalletResponse struct {
		ID string `json:"id"`
	} //@name CreateWalletResponse

	GetBalanceResponse struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	} //@name GetBalanceResponse
)
