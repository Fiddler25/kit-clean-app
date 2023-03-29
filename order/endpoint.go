package order

import (
	"clean-architecture-sample/product"
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	placeOrderRequest struct {
		ProductID product.ID `json:"product_id"`
		UserID    uint32     `json:"user_id"`
		Quantity  uint8      `json:"quantity"`
	}

	placeOrderResponse struct {
		ID         ID         `json:"id,omitempty"`
		ProductID  product.ID `json:"product_id,omitempty"`
		UserID     uint32     `json:"user_id,omitempty"`
		Quantity   uint8      `json:"quantity,omitempty"`
		TotalPrice float64    `json:"total_price,omitempty"`
		Err        error      `json:"err,omitempty" json:"err,omitempty"`
	}
)

func (r placeOrderResponse) error() error { return r.Err }

func makePlaceOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(placeOrderRequest)

		ipt := &placeOrderInput{
			productID: req.ProductID,
			userID:    req.UserID,
			quantity:  req.Quantity,
		}
		opt, err := s.PlaceOrder(ctx, ipt)

		return placeOrderResponse{
			ID:         opt.ID,
			ProductID:  opt.ProductID,
			UserID:     opt.UserID,
			Quantity:   opt.Quantity,
			TotalPrice: opt.TotalPrice,
			Err:        err,
		}, nil
	}
}
