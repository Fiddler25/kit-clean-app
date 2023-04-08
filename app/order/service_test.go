package order

import (
	"context"
	"errors"
	"kit-clean-app/app/model"
	"kit-clean-app/app/product"
	"kit-clean-app/pkg/test"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestService_PlaceOrder(t *testing.T) {
	t.Parallel()

	tx := test.Tx()

	type (
		give struct {
			ipt         *placeOrderInput
			productRepo product.MockRepository
			orderRepo   MockRepository
		}

		want struct {
			order *model.Order
			err   error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			"正常終了",
			give{
				ipt: &placeOrderInput{
					productID: 1,
					userID:    1,
					quantity:  2,
				},
				productRepo: product.MockRepository{
					GetFunc: func(ctx context.Context, id model.ProductID) (*model.Product, error) {
						return &model.Product{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
							Stock:       5,
						}, nil
					},
					UpdateFunc: func(ctx context.Context, p *model.Product) (*model.Product, error) {
						return &model.Product{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
							Stock:       3,
						}, nil
					},
				},
				orderRepo: MockRepository{
					CreateFunc: func(ctx context.Context, e *model.Order) (*model.Order, error) {
						return &model.Order{
							ID:         1,
							ProductID:  1,
							UserID:     1,
							Quantity:   3,
							TotalPrice: 3000,
						}, nil
					},
				},
			},
			want{
				order: &model.Order{
					ID:         1,
					ProductID:  1,
					UserID:     1,
					Quantity:   3,
					TotalPrice: 3000,
				},
			},
		},
		{
			"productRepo.Update()でエラー発生",
			give{
				ipt: &placeOrderInput{
					productID: 1,
					userID:    1,
					quantity:  2,
				},
				productRepo: product.MockRepository{
					GetFunc: func(ctx context.Context, id model.ProductID) (*model.Product, error) {
						return &model.Product{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
							Stock:       5,
						}, nil
					},
					UpdateFunc: func(ctx context.Context, p *model.Product) (*model.Product, error) {
						return &model.Product{}, test.ErrDummy
					},
				},
			},
			want{
				order: &model.Order{},
				err:   test.ErrDummy,
			},
		},
		// ...
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tx, tt.give.orderRepo, tt.give.productRepo)

			got, err := s.PlaceOrder(context.Background(), tt.give.ipt)

			if diff := cmp.Diff(tt.want.order, got); diff != "" {
				t.Errorf("PlaceOrder() mismatch (-want +got)\n%s", diff)
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("err = %v, want = %v", err, tt.want.err)
			}
		})
	}
}