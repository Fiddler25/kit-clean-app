package product

import (
	"context"
	"kit-clean-app/app/model"
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

func (l *loggingService) CreateProduct(ctx context.Context, ipt createProductInput) (p *model.Product, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "create",
			"name", ipt.Name,
			"description", ipt.Description,
			"price", ipt.Price,
			"stock", ipt.Stock,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return l.Service.CreateProduct(ctx, ipt)
}
