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
