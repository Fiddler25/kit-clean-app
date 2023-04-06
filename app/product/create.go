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

func (s *service) CreateProduct(ctx context.Context, ipt createProductInput) (*model.Product, error) {
	m := &model.Product{
		Name:        ipt.Name,
		Description: ipt.Description,
		Price:       ipt.Price,
		Stock:       ipt.Stock,
	}
	p, err := s.repo.Create(ctx, m)
	if err != nil {
		return &model.Product{}, err
	}

	return p, nil
}
