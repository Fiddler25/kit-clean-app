package order

import (
	"context"
	"kit-clean-app/app/model"
	"kit-clean-app/app/product"
	"kit-clean-app/db"
)

type Service interface {
	PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*ReadOrder, error)
}

type service struct {
	tx           db.Tx
	orderStore   Store
	productStore product.Store
}

func NewService(tx db.Tx, orderStore Store, productStore product.Store) Service {
	return &service{
		tx:           tx,
		orderStore:   orderStore,
		productStore: productStore,
	}
}

type ReadOrder struct {
	ID         model.OrderID   `json:"id,omitempty"`
	ProductID  model.ProductID `json:"product_id,omitempty"`
	UserID     uint32          `json:"user_id,omitempty"`
	Quantity   uint8           `json:"quantity,omitempty"`
	TotalPrice float64         `json:"total_price,omitempty"`
}

func modelToRead(m *model.Order) *ReadOrder {
	return &ReadOrder{
		ID:         m.ID,
		ProductID:  m.ProductID,
		UserID:     m.UserID,
		Quantity:   m.Quantity,
		TotalPrice: m.TotalPrice,
	}
}
