// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	port "github.com/mkaiho/go-aws-sandbox/usecase/port"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *UserRepository) Create(ctx context.Context, input port.UserCreateInput) (*port.UserCreateOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *port.UserCreateOutput
	if rf, ok := ret.Get(0).(func(context.Context, port.UserCreateInput) *port.UserCreateOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.UserCreateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, port.UserCreateInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, input
func (_m *UserRepository) DeleteByID(ctx context.Context, input port.UserDeleteByIDInput) (*port.UserDeleteByIDOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *port.UserDeleteByIDOutput
	if rf, ok := ret.Get(0).(func(context.Context, port.UserDeleteByIDInput) *port.UserDeleteByIDOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.UserDeleteByIDOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, port.UserDeleteByIDInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, input
func (_m *UserRepository) FindByEmail(ctx context.Context, input port.UserFindByEmailInput) (*port.UserFindByEmailOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *port.UserFindByEmailOutput
	if rf, ok := ret.Get(0).(func(context.Context, port.UserFindByEmailInput) *port.UserFindByEmailOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.UserFindByEmailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, port.UserFindByEmailInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, input
func (_m *UserRepository) FindByID(ctx context.Context, input port.UserFindByIDInput) (*port.UserFindByIDOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *port.UserFindByIDOutput
	if rf, ok := ret.Get(0).(func(context.Context, port.UserFindByIDInput) *port.UserFindByIDOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.UserFindByIDOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, port.UserFindByIDInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, input
func (_m *UserRepository) List(ctx context.Context, input port.UserListInput) (*port.UserListOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *port.UserListOutput
	if rf, ok := ret.Get(0).(func(context.Context, port.UserListInput) *port.UserListOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.UserListOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, port.UserListInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
