package order

import (
	"context"
	"kit-clean-app/app/model"
)

var _ Service = MockService{}

type MockService struct {
	PlaceOrderFunc func(ctx context.Context, ipt *placeOrderInput) (*ReadOrder, error)
}

func (m MockService) PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*ReadOrder, error) {
	return m.PlaceOrderFunc(ctx, ipt)
}

var _ Store = MockStore{}

type MockStore struct {
	CreateFunc func(ctx context.Context, e *model.Order) (*model.Order, error)
}

func (m MockStore) Create(ctx context.Context, e *model.Order) (*model.Order, error) {
	return m.CreateFunc(ctx, e)
}
