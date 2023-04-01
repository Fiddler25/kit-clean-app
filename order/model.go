package order

import "clean-architecture-sample/product"

type ID uint32

type Order struct {
	ID         ID
	ProductID  product.ID
	UserID     uint32
	Quantity   uint8
	TotalPrice float64
}

func (o *Order) CalcTotalPrice(price float64) {
	o.TotalPrice = price * float64(o.Quantity)
}
