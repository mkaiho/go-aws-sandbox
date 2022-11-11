package adapter

import (
	"context"

	"github.com/mkaiho/go-aws-sandbox/adapter/rdb"
	rdbadapter "github.com/mkaiho/go-aws-sandbox/adapter/rdb"
	"github.com/mkaiho/go-aws-sandbox/usecase/port"
)

var _ port.TxRepository[any] = (*TxManager[any])(nil)

type TxManager[T any] struct {
	rdb rdbadapter.DB
}

func NewTxManager[T any](rdb rdbadapter.DB) *TxManager[T] {
	return &TxManager[T]{
		rdb: rdb,
	}
}

func (tm *TxManager[T]) ContextWithNewTx(ctx context.Context) (context.Context, error) {
	return rdb.ContextWithTx(ctx, tm.rdb)
}

func (tm *TxManager[T]) DoInTx(ctx context.Context, f func(ctx context.Context) (*T, error)) (res *T, err error) {
	var tx rdbadapter.Transaction
	defer func() {
		if tx != nil {
			if err != nil {
				tx.Rollback()
				return
			}
			if cErr := tx.Commit(); cErr != nil {
				err = cErr
			}
		}
	}()

	tx, err = rdb.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return f(ctx)
}
