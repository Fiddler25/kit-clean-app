package product

import (
	"context"
	"encoding/json"
	"errors"
	"kit-clean-app/app/model"
	"kit-clean-app/pkg/apperr"
	"net/http"
	"strconv"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := mux.NewRouter()

	r.Methods(http.MethodPost).Path("/v1/products").Handler(kithttp.NewServer(
		makeCreateProductEndpoint(s),
		decodeCreateProductRequest,
		encodeResponse,
		opts...,
	))
	r.Methods(http.MethodGet).Path("/v1/products/{id}/convert-currency").Handler(kithttp.NewServer(
		makeConvertCurrencyEndpoint(s),
		decodeConvertCurrencyRequest,
		encodeResponse,
		opts...,
	))

	return r
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body createProductRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeConvertCurrencyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, apperr.ErrBadRoute
	}
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}

	currencyCode := r.URL.Query().Get("currency_code")
	if currencyCode == "" {
		return nil, errors.New("currency_code query parameter is required")
	}

	return convertCurrencyRequest{
		ID:           model.ProductID(uint32(productID)),
		CurrencyCode: currencyCode,
	}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch {
	case errors.Is(err, apperr.ErrInvalidArgument):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
