// Code generated by mockery v2.51.1. DO NOT EDIT.

package models

import mock "github.com/stretchr/testify/mock"

// IDTO is an autogenerated mock type for the IDTO type
type IDTO struct {
	mock.Mock
}

type IDTO_Expecter struct {
	mock *mock.Mock
}

func (_m *IDTO) EXPECT() *IDTO_Expecter {
	return &IDTO_Expecter{mock: &_m.Mock}
}

// GetID provides a mock function with no fields
func (_m *IDTO) GetID() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetID")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// IDTO_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type IDTO_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
func (_e *IDTO_Expecter) GetID() *IDTO_GetID_Call {
	return &IDTO_GetID_Call{Call: _e.mock.On("GetID")}
}

func (_c *IDTO_GetID_Call) Run(run func()) *IDTO_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IDTO_GetID_Call) Return(_a0 uint) *IDTO_GetID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IDTO_GetID_Call) RunAndReturn(run func() uint) *IDTO_GetID_Call {
	_c.Call.Return(run)
	return _c
}

// NewIDTO creates a new instance of IDTO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDTO(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDTO {
	mock := &IDTO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
