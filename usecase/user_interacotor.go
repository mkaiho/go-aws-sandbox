package usecase

import (
	"context"
	"errors"

	"github.com/mkaiho/go-aws-sandbox/entity"
	"github.com/mkaiho/go-aws-sandbox/usecase/port"
)

var (
	ErrNotFoundUser  = NewNotFoundError("user")
	ErrDuplicateUser = NewDuplicateError("user")
)

var _ UserInteractor = (*userInteractorImpl)(nil)

func NewUserInteractor(userRepository port.UserRepository) *userInteractorImpl {
	return &userInteractorImpl{
		userRepository: userRepository,
	}
}

type UserInteractor interface {
	Register(ctx context.Context, input UserRegisterInput) (*UserRegisterOutput, error)
}

type (
	UserDetail struct {
		ID    entity.UserID
		Name  string
		Email string
	}
	UserRegisterInput struct {
		Name  string
		Email string
	}
	UserRegisterOutput struct {
		RegisteredUserDetail UserDetail
	}
)

type userInteractorImpl struct {
	userRepository port.UserRepository
}

func (u *userInteractorImpl) Register(ctx context.Context, input UserRegisterInput) (*UserRegisterOutput, error) {
	foundOutput, err := u.userRepository.FindByEmail(ctx, port.UserFindByEmailInput{
		Email: input.Email,
	})
	if err != nil {
		var notFoundError *NotFoundEntityError
		if !errors.As(err, &notFoundError) {
			return nil, err
		}
	}
	if foundOutput != nil {
		return nil, ErrDuplicateUser
	}

	createdUserOutput, err := u.userRepository.Create(ctx, port.UserCreateInput{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}
	createdUser := createdUserOutput.CreatedUser

	return &UserRegisterOutput{
		RegisteredUserDetail: UserDetail{
			ID:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
		},
	}, nil
}
