// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	usecase "github.com/mkaiho/go-aws-sandbox/usecase"
	mock "github.com/stretchr/testify/mock"
)

// UserInteractor is an autogenerated mock type for the UserInteractor type
type UserInteractor struct {
	mock.Mock
}

// Register provides a mock function with given fields: ctx, input
func (_m *UserInteractor) Register(ctx context.Context, input usecase.UserRegisterInput) (*usecase.UserRegisterOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *usecase.UserRegisterOutput
	if rf, ok := ret.Get(0).(func(context.Context, usecase.UserRegisterInput) *usecase.UserRegisterOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.UserRegisterOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, usecase.UserRegisterInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserInteractor interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserInteractor creates a new instance of UserInteractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserInteractor(t mockConstructorTestingTNewUserInteractor) *UserInteractor {
	mock := &UserInteractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}