package order

import (
	"context"
	"kit-clean-app/app/model"
	"kit-clean-app/app/product"
	"kit-clean-app/db"
)

type Service interface {
	PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*model.Order, error)
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
