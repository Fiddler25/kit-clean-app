package product

import (
	"context"
	"kit-clean-app/app/model"
)

type convertCurrencyInput struct {
	id           model.ProductID
	currencyCode string
}

func (s *service) ConvertCurrency(ctx context.Context, ipt convertCurrencyInput) (*ReadProduct, error) {
	return &ReadProduct{
		ID:           100,
		Name:         "",
		Description:  "",
		Price:        0,
		Stock:        0,
		CurrencyCode: "JP",
	}, nil
}
