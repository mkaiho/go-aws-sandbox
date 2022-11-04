package port

import (
	"context"

	"github.com/mkaiho/go-aws-sandbox/entity"
)

type (
	UserListInput struct {
		Limit *uint
	}
	UserListOutput struct {
		User entity.User
	}

	UserFindByIDInput struct {
		ID entity.UserID
	}
	UserFindByIDOutput struct {
		User entity.User
	}

	UserCreateInput struct {
		Name  string
		Email string
	}
	UserCreateOutput struct {
		CreatedUser entity.User
	}

	UserDeleteByIDInput struct {
		ID entity.UserID
	}
	UserDeleteByIDOutput struct {
		User entity.User
	}
)

type UserRepository interface {
	GenerateID() entity.UserID
	ParseID(v string) entity.UserID
	List(ctx context.Context, input UserListInput) (*UserListOutput, error)
	FindByID(ctx context.Context, input UserFindByIDInput) (*UserFindByIDOutput, error)
	Create(ctx context.Context, input UserCreateInput) (*UserCreateOutput, error)
	DeleteByID(ctx context.Context, input UserDeleteByIDInput) (*UserDeleteByIDOutput, error)
}
