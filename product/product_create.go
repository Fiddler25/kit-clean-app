package product

import (
	"context"
)

type createProductInput struct {
	Name        string
	Description string
	Price       float64
}

func (s *service) CreateProduct(ctx context.Context, ipt createProductInput) (*Product, error) {
	e := &Product{
		Name:        ipt.Name,
		Description: ipt.Description,
		Price:       ipt.Price,
	}
	p, err := s.repo.Create(ctx, e)
	if err != nil {
		return &Product{}, err
	}

	return p, nil
}
