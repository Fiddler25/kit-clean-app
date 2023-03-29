package order

import (
	"clean-architecture-sample/db"
	"clean-architecture-sample/product"
	"context"
)

func (r *repository) Create(ctx context.Context, o *Order) (*Order, error) {
	e, err := db.Client(ctx).Order.
		Create().
		SetProductID(uint32(o.ProductID)).
		SetUserID(int(o.UserID)).
		SetQuantity(o.Quantity).
		SetTotalPrice(o.TotalPrice).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &Order{
		ID:         ID(e.ID),
		ProductID:  product.ID(e.ProductID),
		UserID:     uint32(e.UserID),
		Quantity:   e.Quantity,
		TotalPrice: e.TotalPrice,
	}, nil
}
