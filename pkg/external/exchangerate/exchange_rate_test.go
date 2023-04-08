package exchangerate_test

import (
	"errors"
	"fmt"
	"kit-clean-app/pkg/external/exchangerate"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAPI_Convert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		want        float64
		err         error
		testHandler func(w http.ResponseWriter, r *http.Request)
	}{
		{
			name: "正常終了",
			want: 0.007567,
			testHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				fmt.Fprint(w, `{
					"date": "2023-04-08",
					"info": {
						"rate": 0.007567,
						"timestamp": 1680942063
					},
					"query": {
						"amount": 1,
						"from": "JPY",
						"to": "USD"
					},
					"result": 0.007567,
					"success": true
				}`)
			},
		},
		{
			name: "Convert()メソッドでエラー発生",
			want: 0,
			err:  exchangerate.ErrConvert,
			testHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(401)
				fmt.Fprint(w, `{"message": "Invalid authentication credentials"}`)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/exchangerates_data/convert", tt.testHandler)
			testServer := httptest.NewServer(mux)
			defer testServer.Close()

			api, err := exchangerate.New(testServer.URL, "dummy-api-key")
			if err != nil {
				t.Fatal(err)
			}

			got, err := api.Convert("USD")

			if !errors.Is(err, tt.err) {
				t.Errorf("want = %v, error = %v", tt.err, err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("rate mismatch (-want +got)\n%s", diff)
			}
		})
	}
}
