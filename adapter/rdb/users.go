package rdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mkaiho/go-aws-sandbox/entity"
)

type UserIDManager struct {
	generator func() string
	validator func(v string) error
}

var userAllColumns = []string{
	"id",
	"name",
	"email",
}

type UserRow struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
type UserRows []*UserRow

type UserDataAccess struct {
	tx *sqlx.Tx
}

func NewUserDataAccess(tx *sqlx.Tx) *UserDataAccess {
	return &UserDataAccess{
		tx: tx,
	}
}

func (a *UserDataAccess) Create(ctx context.Context, row *UserRow) error {
	q := `INSERT INTO users (id, name, email) VALUES (:id, :name, :email)`
	_, err := a.tx.NamedExecContext(ctx, q, row)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserDataAccess) List(ctx context.Context, limit *uint) (UserRows, error) {
	cols := strings.Join(userAllColumns, ", ")
	q := fmt.Sprintf(`SELECT %s FROM users`, cols)
	var params []interface{}

	if limit != nil {
		q += ` LIMIT = ?`
		params = append(params, *limit)
	}

	var rows UserRows
	err := a.tx.SelectContext(ctx, &rows, q, params...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (a *UserDataAccess) GetByID(ctx context.Context, id entity.UserID) (*UserRow, error) {
	cols := strings.Join(userAllColumns, ", ")
	q := fmt.Sprintf(`SELECT %s FROM users WHERE id = ?`, cols)
	var row UserRow
	err := a.tx.GetContext(ctx, &row, q, id.String())
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (a *UserDataAccess) GetByEmail(ctx context.Context, email string) (*UserRow, error) {
	cols := strings.Join(userAllColumns, ", ")
	q := fmt.Sprintf(`SELECT %s FROM users WHERE email = ?`, cols)
	var row UserRow
	err := a.tx.GetContext(ctx, &row, q, email)
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (a *UserDataAccess) DeleteByID(ctx context.Context, id entity.UserID) error {
	q := `DELETE FROM users WHERE id = ?`
	_, err := a.tx.ExecContext(ctx, q, id.String())
	if err != nil {
		return err
	}

	return nil
}
