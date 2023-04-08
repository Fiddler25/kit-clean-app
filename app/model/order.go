package model

type OrderID uint32

type Order struct {
	ID         OrderID
	ProductID  ProductID
	UserID     uint32
	Quantity   uint8
	TotalPrice float64
}

func (o *Order) CalcTotalPrice(price float64) {
	o.TotalPrice = price * float64(o.Quantity)
}
