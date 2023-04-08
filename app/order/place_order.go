package order

import (
	"context"
	"kit-clean-app/app/model"
)

type placeOrderInput struct {
	productID model.ProductID
	userID    uint32
	quantity  uint8
}

func (s service) PlaceOrder(ctx context.Context, ipt *placeOrderInput) (*ReadOrder, error) {
	curr, err := s.productRepo.Get(ctx, ipt.productID)
	if err != nil {
		return nil, err
	}

	if err := curr.ReduceStock(ipt.quantity); err != nil {
		return &ReadOrder{}, err
	}

	var order = &model.Order{
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
		return &ReadOrder{}, err
	}

	return toRead(order), nil
}
