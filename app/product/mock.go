package product

import (
	"context"
	"kit-clean-app/app/model"
)

var _ Service = MockService{}

type MockService struct {
	CreateProductFunc   func(ctx context.Context, ipt createProductInput) (*model.Product, error)
	ConvertCurrencyFunc func(ctx context.Context, ipt convertCurrencyInput) (*ReadProduct, error)
}

func (m MockService) CreateProduct(ctx context.Context, ipt createProductInput) (*model.Product, error) {
	return m.CreateProductFunc(ctx, ipt)
}

func (m MockService) ConvertCurrency(ctx context.Context, ipt convertCurrencyInput) (*ReadProduct, error) {
	return m.ConvertCurrencyFunc(ctx, ipt)
}

var _ Repository = MockRepository{}

type MockRepository struct {
	CreateFunc func(ctx context.Context, p *model.Product) (*model.Product, error)
	GetFunc    func(ctx context.Context, id model.ProductID) (*model.Product, error)
	UpdateFunc func(ctx context.Context, p *model.Product) (*model.Product, error)
}

func (m MockRepository) Create(ctx context.Context, p *model.Product) (*model.Product, error) {
	return m.CreateFunc(ctx, p)
}

func (m MockRepository) Get(ctx context.Context, id model.ProductID) (*model.Product, error) {
	return m.GetFunc(ctx, id)
}

func (m MockRepository) Update(ctx context.Context, p *model.Product) (*model.Product, error) {
	return m.UpdateFunc(ctx, p)
}
