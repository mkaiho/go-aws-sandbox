package rdb

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mkaiho/go-aws-sandbox/usecase/port"
)

var ctxTxKey = struct{}{}

func ContextWithTx(ctx context.Context, db DB) (context.Context, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, ctxTxKey, tx), nil
}

func TxFromContext(ctx context.Context) (Transaction, error) {
	tx, ok := ctx.Value(ctxTxKey).(Transaction)
	if !ok {
		return nil, errors.New("invalid transaction")
	}
	if tx == nil {
		return nil, port.ErrNoTransaction
	}
	return tx, nil
}

var ErrInvalidDriverName = errors.New("driver name is unknown")

type DriverName string

func (n DriverName) String() string {
	return string(n)
}

type DB interface {
	Begin() (Transaction, error)
}

type Config interface {
	GetDSN() string
	GetMaxConns() int
	GetDriverName() DriverName
}

type Transaction interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	Commit() error
	Rollback() error
}
