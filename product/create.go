package product

import (
	"context"
)

type createProductInput struct {
	Name        string
	Description string
	Price       float64
	Stock       uint8
}

func (s *service) CreateProduct(ctx context.Context, ipt createProductInput) (*Product, error) {
	e := &Product{
		Name:        ipt.Name,
		Description: ipt.Description,
		Price:       ipt.Price,
		Stock:       ipt.Stock,
	}
	p, err := s.repo.Create(ctx, e)
	if err != nil {
		return &Product{}, err
	}

	return p, nil
}
