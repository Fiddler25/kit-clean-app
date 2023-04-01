package order

import (
	"clean-architecture-sample/product"
	"context"
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

	if err := curr.ReduceStock(ipt.quantity); err != nil {
		return &Order{}, err
	}

	var order = &Order{
		ProductID: ipt.productID,
		UserID:    ipt.userID,
		Quantity:  ipt.quantity,
	}
	if err := s.tx.Do(ctx, func(ctx context.Context) error {

		p, err := s.productRepo.Update(ctx, curr)
		if err != nil {
			return err
		}

		order.CalcTotalPrice(p.Price)

		o, err := s.repo.Create(ctx, order)
		if err != nil {
			return err
		}
		order = o

		return nil

	}); err != nil {
		return &Order{}, err
	}

	return &Order{
		ID:         order.ID,
		ProductID:  order.ProductID,
		UserID:     order.UserID,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
	}, nil
}
