package product_test

import (
	"clean-architecture-sample/product"
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestMakeCreateProductEndpoint(t *testing.T) {
	t.Parallel()

	svc := &product.MockService{}

	type give struct {
		req product.ExportCreateProductRequest
		f   product.MockService
	}

	tests := []struct {
		name string
		give give
		want product.ExportCreateProductResponse
	}{
		{
			"【OK】正常系",
			give{
				req: product.ExportCreateProductRequest{
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
				},
				f: product.MockService{
					CreateProductFunc: func(ctx context.Context) (*product.Product, error) {
						return &product.Product{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
						}, nil
					},
				},
			},
			product.ExportCreateProductResponse{
				ID:          1,
				Name:        "コーヒー",
				Description: "豆 深煎り 200g",
				Price:       1500,
			},
		},
		{
			"【NG】nameが空文字",
			give{
				req: product.ExportCreateProductRequest{
					Name:        "",
					Description: "豆 深煎り 200g",
					Price:       1500,
				},
			},
			product.ExportCreateProductResponse{
				Err: product.ErrInvalidArgument,
			},
		},
		{
			"【NG】priceが0未満",
			give{
				req: product.ExportCreateProductRequest{
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       -1,
				},
			},
			product.ExportCreateProductResponse{
				Err: product.ErrInvalidArgument,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			svc.CreateProductFunc = tt.give.f.CreateProductFunc

			resp, _ := product.ExportMakeCreateProductEndpoint(svc)(context.Background(), tt.give.req)
			got, ok := resp.(product.ExportCreateProductResponse)
			if !ok {
				t.Errorf("unexpected response = %v", resp)
			}

			opt := cmpopts.IgnoreFields(product.ExportCreateProductResponse{}, "Err")
			if diff := cmp.Diff(tt.want, got, opt); diff != "" {
				t.Errorf("response mismatch (-want +got)\n%s", diff)
			}

			if !errors.Is(got.Err, tt.want.Err) {
				t.Errorf("err = %v, want = %v", got.Err, tt.want.Err)
			}
		})
	}
}
