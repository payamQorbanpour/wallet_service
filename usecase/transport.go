package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	"wallet_service/usecase/dto"

	"github.com/gorilla/mux"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateWalletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetBalanceRequest
	vars := mux.Vars(r)

	req = dto.GetBalanceRequest{
		ID: vars["id"],
	}

	return req, nil
}

func decodeChargeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.ChargeWalletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
