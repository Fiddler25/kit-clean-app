package product

import (
	"context"
	"errors"
	"kit-clean-app/pkg/apperr"
	"kit-clean-app/pkg/test"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestMakeCreateProductEndpoint(t *testing.T) {
	t.Parallel()

	type give struct {
		req createProductRequest
		svc MockService
	}

	tests := []struct {
		name string
		give give
		want createProductResponse
	}{
		{
			"正常終了",
			give{
				req: createProductRequest{
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
					Stock:       5,
				},
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt createProductInput) (*ReadProduct, error) {
						return &ReadProduct{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
							Stock:       5,
						}, nil
					},
				},
			},
			createProductResponse{
				Product: &ReadProduct{
					ID:          1,
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
					Stock:       5,
				},
			},
		},
		{
			"nameが空文字",
			give{
				req: createProductRequest{
					Name:        "",
					Description: "豆 深煎り 200g",
					Price:       1500,
					Stock:       5,
				},
			},
			createProductResponse{
				Err: apperr.ErrInvalidArgument,
			},
		},
		{
			"priceが0未満",
			give{
				req: createProductRequest{
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       -1,
					Stock:       5,
				},
			},
			createProductResponse{
				Err: apperr.ErrInvalidArgument,
			},
		},
		{
			"その他のエラー",
			give{
				req: createProductRequest{
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
					Stock:       5,
				},
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt createProductInput) (*ReadProduct, error) {
						return &ReadProduct{}, test.ErrDummy
					},
				},
			},
			createProductResponse{
				Product: &ReadProduct{},
				Err:     test.ErrDummy,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			resp, _ := makeCreateProductEndpoint(tt.give.svc)(context.Background(), tt.give.req)
			got, ok := resp.(createProductResponse)
			if !ok {
				t.Errorf("unexpected response = %v", resp)
			}

			opt := cmpopts.IgnoreFields(createProductResponse{}, "Err")
			if diff := cmp.Diff(tt.want, got, opt); diff != "" {
				t.Errorf("response mismatch (-want +got)\n%s", diff)
			}

			if !errors.Is(got.Err, tt.want.Err) {
				t.Errorf("err = %v, want = %v", got.Err, tt.want.Err)
			}
		})
	}
}
