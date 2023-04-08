package product_test

import (
	"context"
	"errors"
	"kit-clean-app/app/model"
	"kit-clean-app/app/product"
	"kit-clean-app/ent"
	"kit-clean-app/ent/enttest"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore_Get(t *testing.T) {
	t.Parallel()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	ctx := ent.NewContext(context.Background(), client)

	store := product.NewStore(client)

	e := &model.Product{
		Name:        "コーヒー",
		Description: "豆 深煎り 200g",
		Price:       1500,
		Stock:       5,
	}
	if _, err := store.Create(ctx, e); err != nil {
		t.Fatal(err)
	}

	type want struct {
		product *model.Product
		err     error
	}

	tests := []struct {
		name string
		id   model.ProductID
		want want
	}{
		{
			"正常終了",
			1,
			want{
				product: &model.Product{
					ID:          1,
					Name:        "コーヒー",
					Description: "豆 深煎り 200g",
					Price:       1500,
					Stock:       5,
				},
			},
		},
		{
			"対象レコードが見つからない",
			999,
			want{
				err: product.ErrNotFound,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.Get(ctx, tt.id)

			if diff := cmp.Diff(tt.want.product, got); diff != "" {
				t.Errorf("product mismatch (-want +got)\n%s", diff)
			}

			if !errors.Is(err, tt.want.err) {
				t.Errorf("err = %v, want = %v", err, tt.want.err)
			}
		})
	}
}
