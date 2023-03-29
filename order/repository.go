package order

import (
	"clean-architecture-sample/ent"
	"context"
)

type Repository interface {
	Create(ctx context.Context, e *Order) (*Order, error)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{client: client}
}

var _ Repository = MockRepository{}

type MockRepository struct {
	CreateFunc func(ctx context.Context, e *Order) (*Order, error)
}

func (m MockRepository) Create(ctx context.Context, e *Order) (*Order, error) {
	return m.CreateFunc(ctx, e)
}
