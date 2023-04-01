package order_test

import (
	"clean-architecture-sample/order"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrder_CalcTotalPrice(t *testing.T) {

	t.Run("【OK】正常系", func(t *testing.T) {
		o := &order.Order{Quantity: 3}
		o.CalcTotalPrice(1000)

		if diff := cmp.Diff(o.TotalPrice, float64(3000)); diff != "" {
			t.Errorf("total price mismatch (-want +got)\n%s", diff)
		}
	})
}
