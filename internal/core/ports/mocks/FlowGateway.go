// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	config "github.com/evertotomalok/text-game/internal/application/config"
	mock "github.com/stretchr/testify/mock"
)

// FlowGateway is an autogenerated mock type for the FlowGateway type
type FlowGateway struct {
	mock.Mock
}

// FlowValidator provides a mock function with given fields: flow
func (_m *FlowGateway) FlowValidator(flow *config.Flow) error {
	ret := _m.Called(flow)

	var r0 error
	if rf, ok := ret.Get(0).(func(*config.Flow) error); ok {
		r0 = rf(flow)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFlowGateway interface {
	mock.TestingT
	Cleanup(func())
}

// NewFlowGateway creates a new instance of FlowGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFlowGateway(t mockConstructorTestingTNewFlowGateway) *FlowGateway {
	mock := &FlowGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
