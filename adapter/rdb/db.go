package rdb

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var ErrInvalidDriverName = errors.New("driver name is unknown")

type DriverName string

func (n DriverName) String() string {
	return string(n)
}

const (
	DriverNameUnknown DriverName = ""
	DriverNameMySQL   DriverName = "mysql"
)

var driverNames = []DriverName{
	DriverNameMySQL,
}

func ParseDriverName(v string) (DriverName, error) {
	parsed := DriverName(v)
	for _, n := range driverNames {
		if parsed == n {
			return n, nil
		}
	}
	return DriverNameUnknown, fmt.Errorf("invalid driver name: %s", v)
}

type Config interface {
	GetDSN() string
	GetMaxConns() int
	GetDriverName() DriverName
}

func Open(conf Config) (*sqlx.DB, error) {
	driverName := conf.GetDriverName()
	if driverName == DriverNameUnknown {
		return nil, ErrInvalidDriverName
	}

	db, err := sqlx.Open(driverName.String(), conf.GetDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(conf.GetMaxConns())

	return db, nil
}
