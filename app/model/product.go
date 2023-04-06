package model

import "errors"

type ProductID uint32

type Product struct {
	ID           ProductID
	Name         string
	Description  string
	Price        float64
	Stock        uint8
	CurrencyCode string
}

var ErrInsufficientStock = errors.New("insufficient stock")

func (p *Product) ReduceStock(quantity uint8) error {
	if p.Stock < quantity {
		return ErrInsufficientStock
	}
	p.Stock -= quantity

	return nil
}
