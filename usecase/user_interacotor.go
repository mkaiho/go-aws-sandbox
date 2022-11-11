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

type (
	UserDetail struct {
		ID    entity.UserID
		Name  string
		Email string
	}
)

// RegisterUserInteractor
var _ RegisterUserInteractor = (*RegisterUserInteractorImpl)(nil)

type RegisterUserInteractor interface {
	Execute(ctx context.Context, input UserRegisterInput) (*UserRegisterOutput, error)
	ExecuteInTx(ctx context.Context, input UserRegisterInput) (*UserRegisterOutput, error)
}

type (
	UserRegisterInput struct {
		Name  string
		Email string
	}
	UserRegisterOutput struct {
		RegisteredUserDetail UserDetail
	}
)

func NewRegisterUserInteractor(
	tx port.TxRepository[UserRegisterOutput],
	userIDManager port.UserIDManager,
	userRepository port.UserRepository,
) *RegisterUserInteractorImpl {
	return &RegisterUserInteractorImpl{
		tx:             tx,
		userIDManager:  userIDManager,
		userRepository: userRepository,
	}
}

type RegisterUserInteractorImpl struct {
	tx             port.TxRepository[UserRegisterOutput]
	userIDManager  port.UserIDManager
	userRepository port.UserRepository
}

func (u *RegisterUserInteractorImpl) Execute(ctx context.Context, input UserRegisterInput) (*UserRegisterOutput, error) {
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

	id, err := u.userIDManager.Generate()
	if err != nil {
		return nil, err
	}

	createdUserOutput, err := u.userRepository.Create(ctx, port.UserCreateInput{
		ID:    id,
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

func (u *RegisterUserInteractorImpl) ExecuteInTx(ctx context.Context, input UserRegisterInput) (*UserRegisterOutput, error) {
	return u.tx.DoInTx(ctx, func(ctx context.Context) (*UserRegisterOutput, error) {
		return u.Execute(ctx, input)
	})
}
