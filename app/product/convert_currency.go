package product

import (
	"context"
	"fmt"
	"kit-clean-app/app/model"
)

type convertCurrencyInput struct {
	id           model.ProductID
	currencyCode string
}

func (s *service) ConvertCurrency(ctx context.Context, ipt convertCurrencyInput) (*ReadProduct, error) {
	p, err := s.repo.Get(ctx, ipt.id)
	if err != nil {
		return nil, err
	}

	rate, err := s.exchangeRateAPI.Convert(ipt.currencyCode)
	if err != nil {
		return nil, err
	}

	p.ConvertPrice(rate)

	return &ReadProduct{
		ID:           100,
		Name:         "",
		Description:  "",
		Price:        0,
		Stock:        0,
		CurrencyCode: "JP",
	}, nil
}
