package product

import (
	"context"
	"kit-clean-app/app/model"
)

type convertCurrencyInput struct {
	id           model.ProductID
	currencyCode string
}

func (s *service) ConvertCurrency(ctx context.Context, ipt *convertCurrencyInput) (*ReadProduct, error) {
	p, err := s.productStore.Get(ctx, ipt.id)
	if err != nil {
		return &ReadProduct{}, err
	}

	rate, err := s.exchangeRateAPI.Convert(ctx, ipt.currencyCode)
	if err != nil {
		return &ReadProduct{}, err
	}

	p.ConvertPrice(rate)

	return modelToRead(p), nil
}
