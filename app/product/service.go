package product

import (
	"context"
	"kit-clean-app/app/model"
	"kit-clean-app/pkg/external/exchangerate"
)

type Service interface {
	CreateProduct(ctx context.Context, ipt createProductInput) (*model.Product, error)
	ConvertCurrency(ctx context.Context, ipt convertCurrencyInput) (*ReadProduct, error)
}

type service struct {
	repo            Repository
	exchangeRateAPI *exchangerate.API
}

func NewService(repo Repository, erAPI *exchangerate.API) Service {
	return &service{
		repo:            repo,
		exchangeRateAPI: erAPI,
	}
}

type ReadProduct struct {
	ID          model.ProductID `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Price       float64         `json:"price,omitempty"`
	Stock       uint8           `json:"stock,omitempty"`
}
