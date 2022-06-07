package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"wallet_service/internal/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/create").Handler(httptransport.NewServer(
		endpoints.CreateWallet,
		decodeCreateRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.GetWallet,
		decodeGetRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/charge").Handler(httptransport.NewServer(
		endpoints.ChargeWallet,
		decodeChargeRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
