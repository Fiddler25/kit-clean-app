package order

import (
	"context"
	"kit-clean-app/app/model"
	"kit-clean-app/db"
	"kit-clean-app/ent"
)

func (r *repository) Create(ctx context.Context, o *model.Order) (*model.Order, error) {
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

func entToOrder(e *ent.Order) *model.Order {
	return &model.Order{
		ID:         model.OrderID(e.ID),
		ProductID:  model.ProductID(e.ProductID),
		UserID:     uint32(e.UserID),
		Quantity:   e.Quantity,
		TotalPrice: e.TotalPrice,
	}
}