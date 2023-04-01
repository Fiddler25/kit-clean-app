package order

import (
	"context"
	"kit-clean-app/db"
	"kit-clean-app/ent"
	"kit-clean-app/product"
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

	return entToOrder(e), nil
}

func entToOrder(e *ent.Order) *Order {
	return &Order{
		ID:         ID(e.ID),
		ProductID:  product.ID(e.ProductID),
		UserID:     uint32(e.UserID),
		Quantity:   e.Quantity,
		TotalPrice: e.TotalPrice,
	}
}
