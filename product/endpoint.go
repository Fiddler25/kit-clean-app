package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

var ErrInvalidArgument = errors.New("invalid argument")

type (
	createProductRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	createProductResponse struct {
		ID          ID      `json:"id,omitempty"`
		Name        string  `json:"name,omitempty"`
		Description string  `json:"description,omitempty"`
		Price       float64 `json:"price,omitempty"`
		Err         error   `json:"error,omitempty"`
	}
)

func (r createProductResponse) error() error { return r.Err }

func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createProductRequest)

		if req.Name == "" {
			return createProductResponse{Err: fmt.Errorf("%w. %s", ErrInvalidArgument, "name is required")}, nil
		}
		if req.Price < 0 {
			return createProductResponse{Err: fmt.Errorf("%w. %s", ErrInvalidArgument, "price is greater than or equal to 0")}, nil
		}

		ipt := createProductInput{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
		}
		opt, err := s.CreateProduct(ctx, ipt)

		return createProductResponse{
			ID:          opt.ID,
			Name:        opt.Name,
			Description: opt.Description,
			Price:       opt.Price,
			Err:         err,
		}, nil
	}
}
