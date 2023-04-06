package product

import (
	"context"
	"kit-clean-app/app/model"
)

type Service interface {
	CreateProduct(ctx context.Context, ipt createProductInput) (*model.Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
