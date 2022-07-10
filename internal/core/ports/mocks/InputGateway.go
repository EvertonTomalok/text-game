// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// InputGateway is an autogenerated mock type for the InputGateway type
type InputGateway struct {
	mock.Mock
}

// GetValidInput provides a mock function with given fields:
func (_m *InputGateway) GetValidInput() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

type mockConstructorTestingTNewInputGateway interface {
	mock.TestingT
	Cleanup(func())
}

// NewInputGateway creates a new instance of InputGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInputGateway(t mockConstructorTestingTNewInputGateway) *InputGateway {
	mock := &InputGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}