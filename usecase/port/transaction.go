package port

import (
	"context"
	"errors"
)

var (
	ErrNoTransaction = errors.New("no transaction")
)

type TxRepository[T any] interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) (*T, error)) (*T, error)
}
