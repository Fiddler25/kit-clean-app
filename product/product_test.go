package product_test

import (
	"clean-architecture-sample/product"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProduct_ReduceStock(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			product  *product.Product
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
			"【OK】正常系",
			give{
				product: &product.Product{
					Stock: 5,
				},
				quantity: 3,
			},
			want{
				stock: 2,
			},
		},
		{
			"【OK】在庫数が注文数と等しい",
			give{
				product: &product.Product{
					Stock: 5,
				},
				quantity: 5,
			},
			want{
				stock: 0,
			},
		},
		{
			"【NG】在庫数が注文数より少ない",
			give{
				product: &product.Product{
					Stock: 5,
				},
				quantity: 6,
			},
			want{
				stock: 5,
				err:   product.ErrInsufficientStock,
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
