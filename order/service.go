package order

import (
	"clean-architecture-sample/db"
	"clean-architecture-sample/product"
	"context"
)

type Service interface {
	PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*Order, error)
}

type service struct {
	tx          db.Tx
	repo        Repository
	productRepo product.Repository
}

func NewService(tx db.Tx, repo Repository, productRepo product.Repository) Service {
	return &service{
		tx:          tx,
		repo:        repo,
		productRepo: productRepo,
	}
}

var _ Service = MockService{}

type MockService struct {
	PlaceOrderFunc func(ctx context.Context, ipt *placeOrderInput) (*Order, error)
}

func (m MockService) PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*Order, error) {
	return m.PlaceOrderFunc(ctx, ipt)
}
