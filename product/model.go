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

func (p *Product) ReduceStock(quantity uint8) error {
	if p.Stock < quantity {
		return errors.New("insufficient stock")
	}
	p.Stock -= quantity

	return nil
}
