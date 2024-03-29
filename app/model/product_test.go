package model_test

import (
	"errors"
	"kit-clean-app/app/model"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProduct_ReduceStock(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			product  *model.Product
			quantity uint8
		}

		want struct {
			stock uint8
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
				product: &model.Product{
					Stock: 5,
				},
				quantity: 3,
			},
			want{
				stock: 2,
			},
		},
		{
			"在庫数が注文数と等しい",
			give{
				product: &model.Product{
					Stock: 5,
				},
				quantity: 5,
			},
			want{
				stock: 0,
			},
		},
		{
			"在庫数が注文数より少ない",
			give{
				product: &model.Product{
					Stock: 5,
				},
				quantity: 6,
			},
			want{
				stock: 5,
				err:   model.ErrInsufficientStock,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.give.product.ReduceStock(tt.give.quantity)

			if diff := cmp.Diff(tt.want.stock, tt.give.product.Stock); diff != "" {
				t.Errorf("stock mismatch (-want +got)\n%s", diff)
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("err = %v, want = %v", err, tt.want.err)
			}
		})
	}
}

func TestProduct_ConvertPrice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		price float64
		rate  float64
		want  float64
	}{
		{"1.5", 100.0, 1.5, 150.0},
		{"0.5", 200.0, 0.5, 100.0},
		{"0", 300.0, 0, 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			product := &model.Product{
				ID:    1,
				Name:  "test",
				Price: tt.price,
				Stock: 5,
			}
			product.ConvertPrice(tt.rate)

			if diff := cmp.Diff(tt.want, product.Price); diff != "" {
				t.Errorf("price mismatch (-want +got)\n%s", diff)
			}
		})
	}
}
