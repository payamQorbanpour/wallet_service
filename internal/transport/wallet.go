package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"wallet_service/internal/dto"

	"github.com/gorilla/mux"
)

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateWalletRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
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

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
