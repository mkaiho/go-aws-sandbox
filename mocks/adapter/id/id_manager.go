// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IDManager is an autogenerated mock type for the IDManager type
type IDManager struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *IDManager) Generate() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: v
func (_m *IDManager) Validate(v string) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIDManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDManager creates a new instance of IDManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDManager(t mockConstructorTestingTNewIDManager) *IDManager {
	mock := &IDManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
