package order

import (
	"clean-architecture-sample/product"
	"context"
	"errors"
)

type placeOrderInput struct {
	productID product.ID
	userID    uint32
	quantity  uint8
}

func (s service) PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*Order, error) {
	curr, err := s.productRepo.Get(ctx, ipt.productID)
	if err != nil {
		return nil, err
	}

	if curr.Stock < ipt.quantity {
		return nil, errors.New("insufficient stock")
	}

	stock := curr.Stock - ipt.quantity
	pe := &product.Product{
		ID:          curr.ID,
		Name:        curr.Name,
		Description: curr.Description,
		Price:       curr.Price,
		Stock:       stock,
	}
	p, err := s.productRepo.Update(ctx, pe)
	if err != nil {
		return nil, err
	}

	price := p.Price * float64(ipt.quantity)
	oe := &Order{
		ProductID:  ipt.productID,
		UserID:     ipt.userID,
		Quantity:   ipt.quantity,
		TotalPrice: price,
	}
	o, err := s.repo.Create(ctx, oe)
	if err != nil {
		return nil, err
	}

	return &Order{
		ID:         o.ID,
		ProductID:  o.ProductID,
		UserID:     o.UserID,
		Quantity:   o.Quantity,
		TotalPrice: o.TotalPrice,
	}, nil
}
