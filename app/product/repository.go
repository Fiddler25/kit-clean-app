package product

import (
	"context"
	"kit-clean-app/app/model"
	"kit-clean-app/ent"
)

type Repository interface {
	Create(ctx context.Context, p *model.Product) (*model.Product, error)
	Get(ctx context.Context, id model.ProductID) (*model.Product, error)
	Update(ctx context.Context, p *model.Product) (*model.Product, error)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{client: client}
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
