package product

import (
	"context"
	"fmt"
	"kit-clean-app/app/model"
	"kit-clean-app/pkg/apperr"

	"github.com/go-kit/kit/endpoint"
)

type (
	createProductRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       uint8   `json:"stock"`
	}

	createProductResponse struct {
		Product *ReadProduct `json:"product,omitempty"`
		Err     error        `json:"error,omitempty"`
	}
)

type (
	convertCurrencyRequest struct {
		ID           model.ProductID `json:"id"`
		CurrencyCode string          `json:"currency_code"`
	}

	convertCurrencyResponse struct {
		Product *ReadProduct `json:"product,omitempty"`
		Err     error        `json:"error,omitempty"`
	}
)

func (r createProductResponse) error() error { return r.Err }

func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createProductRequest)

		if req.Name == "" {
			return createProductResponse{Err: fmt.Errorf("%w. %s", apperr.ErrInvalidArgument, "name is required")}, nil
		}
		if req.Price < 0 {
			return createProductResponse{Err: fmt.Errorf("%w. %s", apperr.ErrInvalidArgument, "price is greater than or equal to 0")}, nil
		}

		ipt := &createProductInput{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Stock:       req.Stock,
		}
		opt, err := s.CreateProduct(ctx, ipt)

		return createProductResponse{
			Product: opt,
			Err:     err,
		}, nil
	}
}

func (r convertCurrencyResponse) error() error { return r.Err }

func makeConvertCurrencyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(convertCurrencyRequest)

		if req.ID == 0 {
			return convertCurrencyResponse{Err: fmt.Errorf("%w. %s", apperr.ErrInvalidArgument, "id is required")}, nil
		}
		if req.CurrencyCode == "" {
			return convertCurrencyResponse{Err: fmt.Errorf("%w. %s", apperr.ErrInvalidArgument, "currency code is required")}, nil
		}

		ipt := &convertCurrencyInput{
			id:           req.ID,
			currencyCode: req.CurrencyCode,
		}
		opt, err := s.ConvertCurrency(ctx, ipt)

		return convertCurrencyResponse{
			Product: opt,
			Err:     err,
		}, nil
	}
}
