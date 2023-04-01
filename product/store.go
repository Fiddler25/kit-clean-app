package product

import (
	"context"
	"errors"
	"kit-clean-app/db"
	"kit-clean-app/ent"
)

var ErrNotFound = errors.New("not found")

func (r *repository) Create(ctx context.Context, p *Product) (*Product, error) {
	e, err := db.Client(ctx).Product.
		Create().
		SetName(p.Name).
		SetDescription(p.Description).
		SetPrice(p.Price).
		SetStock(p.Stock).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToProduct(e), nil
}

func (r *repository) Get(ctx context.Context, id ID) (*Product, error) {
	e, err := db.Client(ctx).Product.Get(ctx, uint32(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return entToProduct(e), nil
}

func (r *repository) Update(ctx context.Context, p *Product) (*Product, error) {
	e, err := db.Client(ctx).Product.
		UpdateOneID(uint32(p.ID)).
		SetName(p.Name).
		SetDescription(p.Description).
		SetPrice(p.Price).
		SetStock(p.Stock).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToProduct(e), nil
}

func entToProduct(e *ent.Product) *Product {
	return &Product{
		ID:          ID(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
	}
}
