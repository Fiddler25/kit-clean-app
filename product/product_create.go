package product

import "context"

type createProductInput struct {
	Name        string
	Description string
	Price       float64
}

func (s *service) CreateProduct(ctx context.Context, ipt createProductInput) (*Product, error) {
	return &Product{
		ID:          1,
		Name:        "コーヒー",
		Description: "豆 深煎り 200g",
		Price:       1500,
	}, nil
}
