package rdb

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

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
