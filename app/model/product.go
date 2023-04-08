package model

import "errors"

type ProductID uint32

type Product struct {
	ID          ProductID
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

func (p *Product) ConvertPrice(rate float64) {
	p.Price = p.Price * rate
}
