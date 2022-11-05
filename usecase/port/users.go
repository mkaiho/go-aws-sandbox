package port

import (
	"context"

	"github.com/mkaiho/go-aws-sandbox/entity"
)

type (
	User struct {
		ID    entity.UserID
		Name  string
		Email string
	}

	UserListInput struct {
		Limit *uint
	}
	UserListOutput struct {
		Users []*User
	}

	UserFindByIDInput struct {
		ID entity.UserID
	}
	UserFindByIDOutput struct {
		User *User
	}

	UserFindByEmailInput struct {
		Email string
	}
	UserFindByEmailOutput struct {
		User *User
	}

	UserCreateInput struct {
		ID    entity.UserID
		Name  string
		Email string
	}
	UserCreateOutput struct {
		CreatedUser *User
	}

	UserDeleteByIDInput struct {
		ID entity.UserID
	}
	UserDeleteByIDOutput struct {
		User *User
	}
)

type UserIDManager interface {
	GenerateID() (entity.UserID, error)
	ParseID(v string) (entity.UserID, error)
}

type UserRepository interface {
	List(ctx context.Context, input UserListInput) (*UserListOutput, error)
	FindByID(ctx context.Context, input UserFindByIDInput) (*UserFindByIDOutput, error)
	FindByEmail(ctx context.Context, input UserFindByEmailInput) (*UserFindByEmailOutput, error)
	Create(ctx context.Context, input UserCreateInput) (*UserCreateOutput, error)
	DeleteByID(ctx context.Context, input UserDeleteByIDInput) (*UserDeleteByIDOutput, error)
}
