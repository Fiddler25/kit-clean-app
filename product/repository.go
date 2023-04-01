package product

import (
	"context"
	"kit-clean-app/ent"
)

type Repository interface {
	Create(ctx context.Context, p *Product) (*Product, error)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{client: client}
}
