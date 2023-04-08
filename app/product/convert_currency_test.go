package product

import (
	"context"
	"errors"
	"kit-clean-app/app/model"
	"kit-clean-app/pkg/external/exchangerate"
	"kit-clean-app/pkg/test"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestService_ConvertCurrency(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			productStore    MockStore
			exchangeRateAPI exchangerate.MockAPI
		}

		want struct {
			product *ReadProduct
			err     error
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
				productStore: MockStore{
					GetFunc: func(ctx context.Context, id model.ProductID) (*model.Product, error) {
						return &model.Product{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
							Stock:       5,
						}, nil
					},
				},
				exchangeRateAPI: exchangerate.MockAPI{
					ConvertFunc: func(ctx context.Context, currencyCode string) (float64, error) {
						return 0.007567, nil
					},
				},
			},
			want{
				product: &ReadProduct{
					ID:          1,
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       11.3505,
					Stock:       5,
				},
				err: nil,
			},
		},
		{
			"productStore.Get()でエラー発生",
			give{
				productStore: MockStore{
					GetFunc: func(ctx context.Context, id model.ProductID) (*model.Product, error) {
						return &model.Product{}, test.ErrDummy
					},
				},
			},
			want{
				product: &ReadProduct{},
				err:     test.ErrDummy,
			},
		},
		{
			"exchangeRateAPI.Convert()でエラー発生",
			give{
				productStore: MockStore{
					GetFunc: func(ctx context.Context, id model.ProductID) (*model.Product, error) {
						return &model.Product{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
							Stock:       5,
						}, nil
					},
				},
				exchangeRateAPI: exchangerate.MockAPI{
					ConvertFunc: func(ctx context.Context, currencyCode string) (float64, error) {
						return 0, test.ErrDummy
					},
				},
			},
			want{
				product: &ReadProduct{},
				err:     test.ErrDummy,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.give.productStore, tt.give.exchangeRateAPI)

			got, err := s.ConvertCurrency(context.Background(), &convertCurrencyInput{id: 1, currencyCode: "USD"})

			if diff := cmp.Diff(tt.want.product, got); diff != "" {
				t.Errorf("ConvertCurrency() mismatch (-want +got)\n%s", diff)
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("err = %v, want = %v", err, tt.want.err)
			}
		})
	}
}
