// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/mkaiho/go-aws-sandbox/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserIDManager is an autogenerated mock type for the UserIDManager type
type UserIDManager struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *UserIDManager) Generate() (entity.UserID, error) {
	ret := _m.Called()

	var r0 entity.UserID
	if rf, ok := ret.Get(0).(func() entity.UserID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entity.UserID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parse provides a mock function with given fields: v
func (_m *UserIDManager) Parse(v string) (entity.UserID, error) {
	ret := _m.Called(v)

	var r0 entity.UserID
	if rf, ok := ret.Get(0).(func(string) entity.UserID); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Get(0).(entity.UserID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(v)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserIDManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserIDManager creates a new instance of UserIDManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserIDManager(t mockConstructorTestingTNewUserIDManager) *UserIDManager {
	mock := &UserIDManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
