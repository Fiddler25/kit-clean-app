package product_test

import (
	"bytes"
	"clean-architecture-sample/product"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/google/go-cmp/cmp"
)

type createProductResponse struct {
	ID          product.ID `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Err         string     `json:"error"`
}

func TestMakeHandler(t *testing.T) {
	t.Parallel()

	m := &product.MockService{}
	logger := log.NewNopLogger()

	type (
		give struct {
			body    map[string]interface{}
			product *product.Product
			err     error
		}

		want struct {
			statusCode int
			resp       createProductResponse
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			"正常系",
			give{
				body: map[string]interface{}{
					"name":        "コーヒー",
					"description": "豆 深煎り 200g",
					"price":       1500,
				},
				product: &product.Product{
					ID:          1,
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
				},
			},
			want{
				statusCode: http.StatusOK,
				resp: createProductResponse{
					ID:          1,
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
				},
			},
		},
		{
			"異常系",
			give{
				body: map[string]interface{}{
					"name":        "コーヒー",
					"description": "豆 深煎り 200g",
					"price":       1500,
				},
				product: &product.Product{},
				err:     errors.New("dummy-error"),
			},
			want{
				statusCode: http.StatusInternalServerError,
				resp: createProductResponse{
					Err: "dummy-error",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m.CreateProductFunc = func(ctx context.Context) (*product.Product, error) {
				return tt.give.product, tt.give.err
			}

			b, err := json.Marshal(tt.give.body)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(b))
			if err != nil {
				t.Fatal(err)
			}

			r := product.MakeHandler(m, logger)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			var got createProductResponse
			if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.want.resp, got); diff != "" {
				t.Errorf("response mismatch (-want +got)\n%s", diff)
			}

			if tt.want.statusCode != resp.StatusCode {
				t.Errorf("status code mismatch want=%d, got=%d", tt.want.statusCode, resp.StatusCode)
			}
		})
	}
}
