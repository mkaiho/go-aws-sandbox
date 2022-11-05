package adapter

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/mkaiho/go-aws-sandbox/adapter/id"
	"github.com/mkaiho/go-aws-sandbox/adapter/rdb"
	"github.com/mkaiho/go-aws-sandbox/entity"
	"github.com/mkaiho/go-aws-sandbox/usecase"
	"github.com/mkaiho/go-aws-sandbox/usecase/port"
)

var (
	_ (port.UserIDManager)  = (*UserIDManager)(nil)
	_ (port.UserRepository) = (*UserGateway)(nil)
)

// UserIDManager
type UserIDManager struct {
	idManager id.IDManager
}

func NewUserIDManager(idManager id.IDManager) *UserIDManager {
	return &UserIDManager{
		idManager: idManager,
	}
}

func (m *UserIDManager) Generate() (entity.UserID, error) {
	id, err := m.idManager.Generate()
	if err != nil {
		return "", err
	}

	return entity.UserID(id), nil
}

func (m *UserIDManager) Parse(v string) (entity.UserID, error) {
	if err := m.idManager.Validate(v); err != nil {
		return "", err
	}

	return entity.UserID(v), nil
}

// UserGateway
type UserGateway struct {
	userIDManager     UserIDManager
	rdbUserDataAccess *rdb.UserDataAccess
}

func NewUserGateway(userIDManager *UserIDManager, tx *sqlx.Tx) *UserGateway {
	return &UserGateway{
		userIDManager:     *userIDManager,
		rdbUserDataAccess: rdb.NewUserDataAccess(tx),
	}
}

func (g *UserGateway) List(ctx context.Context, input port.UserListInput) (*port.UserListOutput, error) {
	rows, err := g.rdbUserDataAccess.List(ctx, input.Limit)
	if err != nil {
		return nil, err
	}

	users := make([]*port.User, len(rows))
	for _, row := range rows {
		id, err := g.userIDManager.Parse(row.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, &port.User{
			ID:    id,
			Name:  row.Name,
			Email: row.Email,
		})
	}

	return &port.UserListOutput{
		Users: users,
	}, nil
}

func (g *UserGateway) FindByID(ctx context.Context, input port.UserFindByIDInput) (*port.UserFindByIDOutput, error) {
	row, err := g.rdbUserDataAccess.GetByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrNotFoundUser
		}
	}
	id, err := g.userIDManager.Parse(row.ID)
	if err != nil {
		return nil, err
	}

	return &port.UserFindByIDOutput{
		User: &port.User{
			ID:    id,
			Name:  row.Name,
			Email: row.Email,
		},
	}, nil
}

func (g *UserGateway) FindByEmail(ctx context.Context, input port.UserFindByEmailInput) (*port.UserFindByEmailOutput, error) {
	row, err := g.rdbUserDataAccess.GetByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrNotFoundUser
		}
	}
	id, err := g.userIDManager.Parse(row.ID)
	if err != nil {
		return nil, err
	}

	return &port.UserFindByEmailOutput{
		User: &port.User{
			ID:    id,
			Name:  row.Name,
			Email: row.Email,
		},
	}, nil
}

func (g *UserGateway) Create(ctx context.Context, input port.UserCreateInput) (*port.UserCreateOutput, error) {
	id, err := g.userIDManager.Generate()
	if err != nil {
		return nil, err
	}

	row := rdb.UserRow{
		ID:    id.String(),
		Name:  input.Name,
		Email: input.Email,
	}
	if err := g.rdbUserDataAccess.Create(ctx, &row); err != nil {
		return nil, err
	}

	return &port.UserCreateOutput{
		CreatedUser: &port.User{
			ID:    id,
			Name:  row.Name,
			Email: row.Email,
		},
	}, nil
}

func (g *UserGateway) DeleteByID(ctx context.Context, input port.UserDeleteByIDInput) (*port.UserDeleteByIDOutput, error) {
	row, err := g.rdbUserDataAccess.GetByID(ctx, input.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrNotFoundUser
		}
	}

	id, err := g.userIDManager.Parse(row.ID)
	if err != nil {
		return nil, err
	}
	if err := g.rdbUserDataAccess.DeleteByID(ctx, id); err != nil {
		return nil, err
	}

	return &port.UserDeleteByIDOutput{
		User: &port.User{
			ID:    id,
			Name:  row.Name,
			Email: row.Email,
		},
	}, nil
}
