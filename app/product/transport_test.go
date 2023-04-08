package product

import (
	"bytes"
	"context"
	"encoding/json"
	"kit-clean-app/pkg/apperr"
	"kit-clean-app/pkg/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/google/go-cmp/cmp"
)

func TestMakeHandler(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			body   map[string]interface{}
			method string
			path   string
			svc    MockService
		}

		want struct {
			statusCode int
			resp       map[string]interface{}
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			"【products】正常終了",
			give{
				body: map[string]interface{}{
					"name":        "コーヒー",
					"description": "豆 深煎り 200g",
					"price":       1500,
					"stock":       5,
				},
				method: http.MethodPost,
				path:   "/v1/products",
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt *createProductInput) (*ReadProduct, error) {
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
			want{
				statusCode: http.StatusOK,
				resp: map[string]interface{}{
					"product": map[string]interface{}{
						"id":          float64(1),
						"name":        "コーヒー",
						"description": "豆 深煎り 200g",
						"price":       float64(1500),
						"stock":       float64(5),
					},
				},
			},
		},
		{
			"【products】不正なリクエスト",
			give{
				body: map[string]interface{}{
					"name":  "コーヒー",
					"price": 1500,
				},
				method: http.MethodPost,
				path:   "/v1/products",
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt *createProductInput) (*ReadProduct, error) {
						return &ReadProduct{}, apperr.ErrInvalidArgument
					},
				},
			},
			want{
				statusCode: http.StatusBadRequest,
				resp: map[string]interface{}{
					"error": "invalid argument",
				},
			},
		},
		{
			"【products】その他のエラー",
			give{
				body: map[string]interface{}{
					"name":  "コーヒー",
					"price": 1500,
				},
				method: http.MethodPost,
				path:   "/v1/products",
				svc: MockService{
					CreateProductFunc: func(ctx context.Context, ipt *createProductInput) (*ReadProduct, error) {
						return &ReadProduct{}, test.ErrDummy
					},
				},
			},
			want{
				statusCode: http.StatusInternalServerError,
				resp: map[string]interface{}{
					"error": "dummy-error",
				},
			},
		},
		{
			"【convert-currency】正常終了",
			give{
				body:   map[string]interface{}{},
				method: http.MethodGet,
				path:   "/v1/products/1/convert-currency?currency_code=USD",
				svc: MockService{
					ConvertCurrencyFunc: func(ctx context.Context, ipt *convertCurrencyInput) (*ReadProduct, error) {
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
			want{
				statusCode: http.StatusOK,
				resp: map[string]interface{}{
					"product": map[string]interface{}{
						"id":          float64(1),
						"name":        "コーヒー",
						"description": "豆 深煎り 200g",
						"price":       float64(1500),
						"stock":       float64(5),
					},
				},
			},
		},
		{
			"【convert-currency】currency_codeが存在しない",
			give{
				body:   map[string]interface{}{},
				method: http.MethodGet,
				path:   "/v1/products/1/convert-currency",
				svc: MockService{
					ConvertCurrencyFunc: func(ctx context.Context, ipt *convertCurrencyInput) (*ReadProduct, error) {
						return &ReadProduct{}, nil
					},
				},
			},
			want{
				statusCode: http.StatusInternalServerError,
				resp: map[string]interface{}{
					"error": "currency_code query parameter is required",
				},
			},
		},
		{
			"【convert-currency】その他のエラー",
			give{
				body:   map[string]interface{}{},
				method: http.MethodGet,
				path:   "/v1/products/1/convert-currency?currency_code=USD",
				svc: MockService{
					ConvertCurrencyFunc: func(ctx context.Context, ipt *convertCurrencyInput) (*ReadProduct, error) {
						return &ReadProduct{}, test.ErrDummy
					},
				},
			},
			want{
				statusCode: http.StatusInternalServerError,
				resp: map[string]interface{}{
					"error": "dummy-error",
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

			req, err := http.NewRequest(tt.give.method, tt.give.path, bytes.NewReader(b))
			if err != nil {
				t.Fatal(err)
			}

			r := MakeHandler(tt.give.svc, log.NewNopLogger())
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			var got map[string]interface{}
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
