package product

import "errors"

type ID uint32

type Product struct {
	ID          ID
	Name        string
	Description string
	Price       float64
	Stock       uint8
}

var ErrInsufficientStock = errors.New("insufficient stock")

func (p *Product) ReduceStock(quantity uint8) error {
	if p.Stock < quantity {
		return ErrInsufficientStock
	}
	p.Stock -= quantity

	return nil
}
