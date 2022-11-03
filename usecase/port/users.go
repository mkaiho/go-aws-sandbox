package port

import (
	"context"

	"github.com/mkaiho/go-aws-sandbox/entity"
)

type (
	UsersListInput struct {
		Limit *uint
	}
	UsersListOutput struct {
		Users entity.Users
	}

	UsersFindByIDInput struct {
		ID entity.UserID
	}
	UsersFindByIDOutput struct {
		User entity.User
	}

	UsersCreateInput struct {
		ID    entity.UserID
		Name  string
		Email string
	}
	UsersCreateOutput struct {
		CreatedUser entity.User
	}

	UsersDeleteByIDInput struct {
		ID entity.UserID
	}
	UsersDeleteByIDOutput struct {
		User entity.User
	}
)

type Users interface {
	GenerateID() entity.UserID
	ParseID() entity.UserID
	List(ctx context.Context, input UsersListInput) (*UsersListOutput, error)
	FindByID(ctx context.Context, input UsersFindByIDInput) (*UsersFindByIDOutput, error)
	Create(ctx context.Context, input UsersCreateInput) (*UsersCreateOutput, error)
	DeleteByID(ctx context.Context, input UsersDeleteByIDInput) (*UsersDeleteByIDOutput, error)
}
