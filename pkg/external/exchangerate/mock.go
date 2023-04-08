package exchangerate

import "context"

var _ API = MockAPI{}

type MockAPI struct {
	ConvertFunc func(ctx context.Context, currencyCode string) (float64, error)
}

func (m MockAPI) Convert(ctx context.Context, currencyCode string) (float64, error) {
	return m.ConvertFunc(ctx, currencyCode)
}
