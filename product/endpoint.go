package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	createProductRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	createProductResponse struct {
		Product *Product `json:"product,omitempty"`
		Err     error    `json:"error,omitempty"`
	}
)

func (r createProductResponse) error() error { return r.Err }

func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createProductRequest)

		ipt := createProductInput{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
		}
		p, err := s.CreateProduct(ctx, ipt)

		return createProductResponse{Product: p, Err: err}, nil
	}
}
