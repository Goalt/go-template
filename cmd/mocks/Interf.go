// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Interf is an autogenerated mock type for the Interf type
type Interf struct {
	mock.Mock
}

// Test provides a mock function with given fields: _a0
func (_m *Interf) Test(_a0 int) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}