package order

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (l *loggingService) PlaceOrder(ctx context.Context, ipt *placeOrderInput) (o *Order, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "create",
			"product_id", ipt.productID,
			"user_id", ipt.userID,
			"quantity", ipt.quantity,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return l.Service.PlaceOrder(ctx, ipt)
}
