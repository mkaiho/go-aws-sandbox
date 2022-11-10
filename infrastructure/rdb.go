package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/mkaiho/go-aws-sandbox/adapter/rdb"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DriverNameUnknown rdb.DriverName = ""
	DriverNameMySQL   rdb.DriverName = "mysql"
)

var driverNames = []rdb.DriverName{
	DriverNameMySQL,
}

func ParseDriverName(v string) (rdb.DriverName, error) {
	parsed := rdb.DriverName(v)
	for _, n := range driverNames {
		if parsed == n {
			return n, nil
		}
	}
	return DriverNameUnknown, fmt.Errorf("invalid driver name: %s", v)
}

// MySQL Configuration
var _ rdb.Config = (*MySQLConfig)(nil)

func LoadMySQLConfig() (*MySQLConfig, error) {
	var c MySQLConfig
	if err := envconfig.Process("MYSQL", &c); err != nil {
		return nil, err
	}
	return &c, nil
}

type MySQLConfig struct {
	Host     string `envconfig:"HOST" required:"true"`
	User     string `envconfig:"USER" required:"true"`
	Database string `envconfig:"DATABASE" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Port     int    `envconfig:"PORT" default:"3306"`
	MaxConns int    `envconfig:"MAX_CONNS" default:"1"`
}

func (c *MySQLConfig) GetDriverName() rdb.DriverName {
	return DriverNameMySQL
}

func (c *MySQLConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4", c.User, c.Password, c.Host, c.Port, c.Database)
}

func (c *MySQLConfig) GetMaxConns() int {
	return c.MaxConns
}

// RDB
var _ rdb.DB = (*RDB)(nil)

type RDB struct {
	db *sqlx.DB
}

func (db *RDB) Begin() (rdb.Transaction, error) {
	tx, err := db.db.Beginx()
	if err != nil {
		return nil, err
	}

	return &RDBTransaction{
		tx: tx,
	}, err
}

func OpenRDB(conf rdb.Config) (*RDB, error) {
	driverName := conf.GetDriverName()
	if driverName == DriverNameUnknown {
		return nil, rdb.ErrInvalidDriverName
	}

	db, err := sqlx.Open(driverName.String(), conf.GetDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(conf.GetMaxConns())

	return &RDB{
		db: db,
	}, nil
}

// RDBTransaction
var _ (rdb.Transaction) = (*RDBTransaction)(nil)

type RDBTransaction struct {
	tx *sqlx.Tx
}

func (rt *RDBTransaction) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return rt.tx.GetContext(ctx, dest, query, args...)
}

func (rt *RDBTransaction) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return rt.tx.SelectContext(ctx, dest, query, args...)
}

func (rt *RDBTransaction) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return rt.tx.NamedExecContext(ctx, query, arg)
}

func (rt *RDBTransaction) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return rt.tx.ExecContext(ctx, query, args...)
}

func (rt *RDBTransaction) Commit() error {
	return rt.tx.Commit()
}

func (rt *RDBTransaction) Rollback() error {
	return rt.tx.Rollback()
}
