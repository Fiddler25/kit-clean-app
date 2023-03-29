package order

import (
	"clean-architecture-sample/product"
	"context"
)

type Service interface {
	PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*Order, error)
}

type service struct {
	repo        Repository
	productRepo product.Repository
}

func NewService(repo Repository, productRepo product.Repository) Service {
	return &service{
		repo:        repo,
		productRepo: productRepo,
	}
}
