package test

import "context"

type MockTx struct {
	doFunc func(context.Context, func(context.Context) error) error
}

func (m MockTx) Do(ctx context.Context, f func(ctx context.Context) error) error {
	return m.doFunc(ctx, f)
}

func Tx() *MockTx {
	return &MockTx{
		doFunc: func(ctx context.Context, f func(context.Context) error) error {
			return f(ctx)
		},
	}
}
