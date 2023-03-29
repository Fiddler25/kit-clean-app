package product

import (
	"clean-architecture-sample/db"
	"context"
)

func (r *repository) Create(ctx context.Context, p *Product) (*Product, error) {
	e, err := r.client.Product.
		Create().
		SetName(p.Name).
		SetDescription(p.Description).
		SetPrice(p.Price).
		SetStock(p.Stock).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:          ID(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
	}, nil
}

func (r *repository) Get(ctx context.Context, id ID) (*Product, error) {
	e, err := r.client.Product.Get(ctx, uint32(id))
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:          ID(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
	}, nil
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

	return &Product{
		ID:          ID(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
	}, nil
}
