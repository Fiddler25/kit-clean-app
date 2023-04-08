package order

import (
	"context"
	"kit-clean-app/app/model"

	"github.com/go-kit/kit/endpoint"
)

type (
	placeOrderRequest struct {
		ProductID model.ProductID `json:"product_id"`
		UserID    uint32          `json:"user_id"`
		Quantity  uint8           `json:"quantity"`
	}

	placeOrderResponse struct {
		Order *ReadOrder `json:"order,omitempty"`
		Err   error      `json:"err,omitempty"`
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
			Order: opt,
			Err:   err,
		}, nil
	}
}
