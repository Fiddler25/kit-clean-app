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
