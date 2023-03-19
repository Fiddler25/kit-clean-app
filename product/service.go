package product

import "context"

type Service interface {
	CreateProduct(ctx context.Context, ipt createProductInput) (*Product, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
