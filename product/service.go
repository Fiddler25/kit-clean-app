package product

import "context"

type Service interface {
	CreateProduct(ctx context.Context, ipt createProductInput) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

var _ Service = MockService{}

type MockService struct {
	CreateProductFunc func(ctx context.Context) (*Product, error)
}

func (m MockService) CreateProduct(ctx context.Context, ipt createProductInput) (*Product, error) {
	return m.CreateProductFunc(ctx)
}
