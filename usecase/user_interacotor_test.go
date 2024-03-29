package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/mkaiho/go-aws-sandbox/entity"
	portmocks "github.com/mkaiho/go-aws-sandbox/mocks/usecase/port"
	"github.com/mkaiho/go-aws-sandbox/usecase/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userInteractorImpl_Execute(t *testing.T) {
	registerdUser := UserDetail{
		ID:    "test_user_id",
		Name:  "test_user_name",
		Email: "test@xxx.com",
	}
	type userIDGenerateMock struct {
		ID  entity.UserID
		err error
	}
	type userFindByEmailMockResult struct {
		output *port.UserFindByEmailOutput
		err    error
	}
	type userCreateMockResult struct {
		output *port.UserCreateOutput
		err    error
	}
	type mocks struct {
		userIDGenerateMock        userIDGenerateMock
		userFindByEmailMockResult userFindByEmailMockResult
		userCreateMockResult      userCreateMockResult
	}
	type args struct {
		ctx   context.Context
		input UserRegisterInput
	}
	tests := []struct {
		name      string
		mocks     mocks
		args      args
		want      *UserRegisterOutput
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "user is registered and returned when no duplicate user exists",
			mocks: mocks{
				userIDGenerateMock: userIDGenerateMock{
					ID:  registerdUser.ID,
					err: nil,
				},
				userFindByEmailMockResult: userFindByEmailMockResult{
					output: nil,
					err:    ErrNotFoundUser,
				},
				userCreateMockResult: userCreateMockResult{
					output: &port.UserCreateOutput{
						CreatedUser: &port.User{
							ID:    registerdUser.ID,
							Name:  registerdUser.Name,
							Email: registerdUser.Email,
						},
					},
					err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: UserRegisterInput{
					Name:  registerdUser.Name,
					Email: registerdUser.Email,
				},
			},
			want: &UserRegisterOutput{
				RegisteredUserDetail: registerdUser,
			},
			assertion: assert.NoError,
		},
		{
			name: "error is returned when user exists",
			mocks: mocks{
				userIDGenerateMock: userIDGenerateMock{
					ID:  registerdUser.ID,
					err: nil,
				},
				userFindByEmailMockResult: userFindByEmailMockResult{
					output: &port.UserFindByEmailOutput{
						User: &port.User{
							ID:    registerdUser.ID,
							Name:  registerdUser.Name,
							Email: registerdUser.Email,
						},
					},
					err: nil,
				},
				userCreateMockResult: userCreateMockResult{
					output: nil,
					err:    nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: UserRegisterInput{
					Name:  registerdUser.Name,
					Email: registerdUser.Email,
				},
			},
			want: nil,
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, ErrDuplicateUser.Error(), "userInteractorImpl.Register() error = %v, wantErr %v", err, ErrDuplicateUser)
			},
		},
		{
			name: "error is returned when failed to get user by email",
			mocks: mocks{
				userIDGenerateMock: userIDGenerateMock{
					ID:  registerdUser.ID,
					err: nil,
				},
				userFindByEmailMockResult: userFindByEmailMockResult{
					output: nil,
					err:    errors.New("failed to get user by email"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: UserRegisterInput{
					Name:  registerdUser.Name,
					Email: registerdUser.Email,
				},
			},
			want: nil,
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				wantErr := errors.New("failed to get user by email")
				return assert.EqualError(t, err, wantErr.Error(), "userInteractorImpl.Register() error = %v, wantErr %v", err, wantErr)
			},
		},
		{
			name: "error is returned when failed to generate user id",
			mocks: mocks{
				userIDGenerateMock: userIDGenerateMock{
					ID:  registerdUser.ID,
					err: errors.New("failed to generate id"),
				},
				userFindByEmailMockResult: userFindByEmailMockResult{
					output: nil,
					err:    ErrNotFoundUser,
				},
			},
			args: args{
				ctx: context.Background(),
				input: UserRegisterInput{
					Name:  registerdUser.Name,
					Email: registerdUser.Email,
				},
			},
			want: nil,
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				wantErr := errors.New("failed to generate id")
				return assert.EqualError(t, err, wantErr.Error(), "userInteractorImpl.Register() error = %v, wantErr %v", err, wantErr)
			},
		},
		{
			name: "error is returned when failed to create user",
			mocks: mocks{
				userIDGenerateMock: userIDGenerateMock{
					ID:  registerdUser.ID,
					err: nil,
				},
				userFindByEmailMockResult: userFindByEmailMockResult{
					output: nil,
					err:    ErrNotFoundUser,
				},
				userCreateMockResult: userCreateMockResult{
					output: nil,
					err:    errors.New("failed to create user"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: UserRegisterInput{
					Name:  registerdUser.Name,
					Email: registerdUser.Email,
				},
			},
			want: nil,
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				wantErr := errors.New("failed to create user")
				return assert.EqualError(t, err, wantErr.Error(), "userInteractorImpl.Register() error = %v, wantErr %v", err, wantErr)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userIDManager := new(portmocks.UserIDManager)
			userIDManager.
				On("Generate").
				Return(tt.mocks.userIDGenerateMock.ID, tt.mocks.userIDGenerateMock.err)
			userRepository := new(portmocks.UserRepository)
			userRepository.
				On("FindByEmail", mock.Anything, port.UserFindByEmailInput{
					Email: tt.args.input.Email,
				}).
				Return(tt.mocks.userFindByEmailMockResult.output, tt.mocks.userFindByEmailMockResult.err)
			userRepository.
				On("Create", mock.Anything, port.UserCreateInput{
					ID:    tt.mocks.userIDGenerateMock.ID,
					Name:  tt.args.input.Name,
					Email: tt.args.input.Email,
				}).
				Return(tt.mocks.userCreateMockResult.output, tt.mocks.userCreateMockResult.err)
			u := &RegisterUserInteractorImpl{
				userIDManager:  userIDManager,
				userRepository: userRepository,
			}
			got, err := u.Execute(tt.args.ctx, tt.args.input)
			if !tt.assertion(t, err) {
				return
			}
			if !assert.Equal(t, got, tt.want, "userInteractorImpl.Register() = %v, want %v", got, tt.want) {
				return
			}
		})
	}
}
