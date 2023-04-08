package model_test

import (
	"kit-clean-app/app/model"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrder_CalcTotalPrice(t *testing.T) {

	t.Run("正常終了", func(t *testing.T) {
		o := &model.Order{Quantity: 3}
		o.CalcTotalPrice(1000)

		if diff := cmp.Diff(o.TotalPrice, float64(3000)); diff != "" {
			t.Errorf("total price mismatch (-want +got)\n%s", diff)
		}
	})
}
