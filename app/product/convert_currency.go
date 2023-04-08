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
	p, err := s.productStore.Get(ctx, ipt.id)
	if err != nil {
		return &ReadProduct{}, err
	}

	rate, err := s.exchangeRateAPI.Convert(ipt.currencyCode)
	if err != nil {
		return &ReadProduct{}, err
	}

	p.ConvertPrice(rate)

	return &ReadProduct{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
	}, nil
}
