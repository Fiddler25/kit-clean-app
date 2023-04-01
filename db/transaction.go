package db

import (
	"context"
	"fmt"
	"kit-clean-app/ent"
)

type Tx interface {
	Do(context.Context, func(context.Context) error) error
}

func (d *DB) Do(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := d.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	txctx := ent.NewTxContext(ctx, tx)

	if err := f(txctx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
