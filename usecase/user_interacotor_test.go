package usecase

import (
	"context"
	"errors"
	"testing"

	portmocks "github.com/mkaiho/go-aws-sandbox/mocks/usecase/port"
	"github.com/mkaiho/go-aws-sandbox/usecase/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userInteractorImpl_Register(t *testing.T) {
	registerdUser := UserDetail{
		ID:    "test_user_id",
		Name:  "test_user_name",
		Email: "test@xxx.com",
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
			name: "error returned when user exists",
			mocks: mocks{
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
			name: "error returned when user cannot be created",
			mocks: mocks{
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
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepository := new(portmocks.UserRepository)
			userRepository.
				On("FindByEmail", mock.Anything, port.UserFindByEmailInput{
					Email: tt.args.input.Email,
				}).
				Return(tt.mocks.userFindByEmailMockResult.output, tt.mocks.userFindByEmailMockResult.err)
			userRepository.
				On("Create", mock.Anything, port.UserCreateInput{
					Name:  tt.args.input.Name,
					Email: tt.args.input.Email,
				}).
				Return(tt.mocks.userCreateMockResult.output, tt.mocks.userCreateMockResult.err)
			u := &userInteractorImpl{
				userRepository: userRepository,
			}
			got, err := u.Register(tt.args.ctx, tt.args.input)
			if !tt.assertion(t, err) {
				return
			}
			if !assert.Equal(t, got, tt.want, "userInteractorImpl.Register() = %v, want %v", got, tt.want) {
				return
			}
		})
	}
}
