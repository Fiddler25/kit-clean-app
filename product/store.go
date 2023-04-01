package product

import "context"

func (r *repository) Create(ctx context.Context, p *Product) (*Product, error) {
	e, err := r.client.Product.
		Create().
		SetName(p.Name).
		SetDescription(p.Description).
		SetPrice(p.Price).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:          ID(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
	}, nil
}
