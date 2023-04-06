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

var _ Service = MockService{}

type MockService struct {
	CreateProductFunc   func(ctx context.Context, ipt createProductInput) (*model.Product, error)
}

func (m MockService) CreateProduct(ctx context.Context, ipt createProductInput) (*model.Product, error) {
	return m.CreateProductFunc(ctx, ipt)
}
