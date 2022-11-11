package rdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/mkaiho/go-aws-sandbox/entity"
)

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

type UserDataAccess struct{}

func NewUserDataAccess() *UserDataAccess {
	return &UserDataAccess{}
}

func (a *UserDataAccess) Create(ctx context.Context, tx Transaction, row *UserRow) error {
	q := `INSERT INTO users (id, name, email) VALUES (:id, :name, :email)`
	_, err := tx.NamedExec(ctx, q, row)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserDataAccess) List(ctx context.Context, tx Transaction, limit *uint) (UserRows, error) {
	cols := strings.Join(userAllColumns, ", ")
	q := fmt.Sprintf(`SELECT %s FROM users`, cols)
	var params []interface{}

	if limit != nil {
		q += ` LIMIT = ?`
		params = append(params, *limit)
	}

	var rows UserRows
	err := tx.Select(ctx, &rows, q, params...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (a *UserDataAccess) GetByID(ctx context.Context, tx Transaction, id entity.UserID) (*UserRow, error) {
	cols := strings.Join(userAllColumns, ", ")
	q := fmt.Sprintf(`SELECT %s FROM users WHERE id = ?`, cols)
	var row UserRow
	err := tx.Get(ctx, &row, q, id.String())
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (a *UserDataAccess) GetByEmail(ctx context.Context, tx Transaction, email string) (*UserRow, error) {
	cols := strings.Join(userAllColumns, ", ")
	q := fmt.Sprintf(`SELECT %s FROM users WHERE email = ?`, cols)
	var row UserRow
	err := tx.Get(ctx, &row, q, email)
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func (a *UserDataAccess) DeleteByID(ctx context.Context, tx Transaction, id entity.UserID) error {
	q := `DELETE FROM users WHERE id = ?`
	_, err := tx.Exec(ctx, q, id.String())
	if err != nil {
		return err
	}

	return nil
}
