package order

import (
	"context"
	"kit-clean-app/app/model"
	"kit-clean-app/ent"
)

type Repository interface {
	Create(ctx context.Context, e *model.Order) (*model.Order, error)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{client: client}
}

var _ Repository = MockRepository{}

type MockRepository struct {
	CreateFunc func(ctx context.Context, e *model.Order) (*model.Order, error)
}

func (m MockRepository) Create(ctx context.Context, e *model.Order) (*model.Order, error) {
	return m.CreateFunc(ctx, e)
}
