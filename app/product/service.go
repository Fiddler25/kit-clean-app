package product

import (
	"context"
	"kit-clean-app/app/model"
	"kit-clean-app/pkg/external/exchangerate"
)

type Service interface {
	CreateProduct(ctx context.Context, ipt createProductInput) (*ReadProduct, error)
	ConvertCurrency(ctx context.Context, ipt convertCurrencyInput) (*ReadProduct, error)
}

type service struct {
	productStore    Store
	exchangeRateAPI *exchangerate.API
}

func NewService(productStore Store, erAPI *exchangerate.API) Service {
	return &service{
		productStore:    productStore,
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

func toRead(m *model.Product) *ReadProduct {
	return &ReadProduct{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Stock:       m.Stock,
	}
}
