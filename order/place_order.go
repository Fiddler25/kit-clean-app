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

	var order *Order
	if err := s.tx.Do(ctx, func(ctx context.Context) error {

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
			return err
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
			return err
		}
		order = o

		return nil

	}); err != nil {
		return nil, err
	}

	return &Order{
		ID:         order.ID,
		ProductID:  order.ProductID,
		UserID:     order.UserID,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
	}, nil
}
