package product

import (
	"context"
	"errors"
	"kit-clean-app/app/model"
	"kit-clean-app/db"
	"kit-clean-app/ent"
)

var ErrNotFound = errors.New("not found")

type Store interface {
	Create(ctx context.Context, p *model.Product) (*model.Product, error)
	Get(ctx context.Context, id model.ProductID) (*model.Product, error)
	Update(ctx context.Context, p *model.Product) (*model.Product, error)
}

type store struct {
	client *ent.Client
}

func NewStore(client *ent.Client) Store {
	return &store{client: client}
}

func (s *store) Create(ctx context.Context, p *model.Product) (*model.Product, error) {
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

func (s *store) Get(ctx context.Context, id model.ProductID) (*model.Product, error) {
	e, err := db.Client(ctx).Product.Get(ctx, uint32(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return entToProduct(e), nil
}

func (s *store) Update(ctx context.Context, p *model.Product) (*model.Product, error) {
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

func entToProduct(e *ent.Product) *model.Product {
	return &model.Product{
		ID:          model.ProductID(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
	}
}
