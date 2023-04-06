package product

import (
	"bytes"
	"context"
	"encoding/json"
	"kit-clean-app/app/model"
	"kit-clean-app/pkg/apperr"
	"kit-clean-app/pkg/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/google/go-cmp/cmp"
)

type testCreateProductResponse struct {
	ID          model.ProductID `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Stock       uint8           `json:"stock"`
	Err         string          `json:"error"`
}

func TestMakeHandler(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			body map[string]interface{}
			svc  MockService
		}

		want struct {
			statusCode int
			resp       testCreateProductResponse
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
				body: map[string]interface{}{
					"name":        "コーヒー",
					"description": "豆 深煎り 200g",
					"price":       1500,
					"stock":       5,
				},
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt createProductInput) (*model.Product, error) {
						return &model.Product{
							ID:          1,
							Name:        "コーヒー",
							Description: "豆 深煎り 200g",
							Price:       1500,
							Stock:       5,
						}, nil
					},
				},
			},
			want{
				statusCode: http.StatusOK,
				resp: testCreateProductResponse{
					ID:          1,
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
					Stock:       5,
				},
			},
		},
		{
			"不正なリクエスト",
			give{
				body: map[string]interface{}{
					"name":  "コーヒー",
					"price": 1500,
				},
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt createProductInput) (*model.Product, error) {
						return &model.Product{}, apperr.ErrInvalidArgument
					},
				},
			},
			want{
				statusCode: http.StatusBadRequest,
				resp: testCreateProductResponse{
					Err: "invalid argument",
				},
			},
		},
		{
			"その他のエラー",
			give{
				body: map[string]interface{}{
					"name":  "コーヒー",
					"price": 1500,
				},
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt createProductInput) (*model.Product, error) {
						return &model.Product{}, test.ErrDummy
					},
				},
			},
			want{
				statusCode: http.StatusInternalServerError,
				resp: testCreateProductResponse{
					Err: "dummy-error",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.give.body)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(b))
			if err != nil {
				t.Fatal(err)
			}

			r := MakeHandler(tt.give.svc, log.NewNopLogger())
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			var got testCreateProductResponse
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
