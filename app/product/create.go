package product

import (
	"context"
	"kit-clean-app/app/model"
)

type createProductInput struct {
	Name        string
	Description string
	Price       float64
	Stock       uint8
}

func (s *service) CreateProduct(ctx context.Context, ipt createProductInput) (*ReadProduct, error) {
	m := &model.Product{
		Name:        ipt.Name,
		Description: ipt.Description,
		Price:       ipt.Price,
		Stock:       ipt.Stock,
	}
	p, err := s.productStore.Create(ctx, m)
	if err != nil {
		return &ReadProduct{}, err
	}

	return modelToRead(p), nil
}
