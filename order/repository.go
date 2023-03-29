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
