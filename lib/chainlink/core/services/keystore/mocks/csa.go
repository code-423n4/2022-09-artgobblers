// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	csakey "github.com/smartcontractkit/chainlink/core/services/keystore/keys/csakey"

	mock "github.com/stretchr/testify/mock"
)

// CSA is an autogenerated mock type for the CSA type
type CSA struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *CSA) Add(key csakey.KeyV2) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(csakey.KeyV2) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *CSA) Create() (csakey.KeyV2, error) {
	ret := _m.Called()

	var r0 csakey.KeyV2
	if rf, ok := ret.Get(0).(func() csakey.KeyV2); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *CSA) Delete(id string) (csakey.KeyV2, error) {
	ret := _m.Called(id)

	var r0 csakey.KeyV2
	if rf, ok := ret.Get(0).(func(string) csakey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields:
func (_m *CSA) EnsureKey() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *CSA) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *CSA) Get(id string) (csakey.KeyV2, error) {
	ret := _m.Called(id)

	var r0 csakey.KeyV2
	if rf, ok := ret.Get(0).(func(string) csakey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *CSA) GetAll() ([]csakey.KeyV2, error) {
	ret := _m.Called()

	var r0 []csakey.KeyV2
	if rf, ok := ret.Get(0).(func() []csakey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]csakey.KeyV2)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetV1KeysAsV2 provides a mock function with given fields:
func (_m *CSA) GetV1KeysAsV2() ([]csakey.KeyV2, error) {
	ret := _m.Called()

	var r0 []csakey.KeyV2
	if rf, ok := ret.Get(0).(func() []csakey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]csakey.KeyV2)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: keyJSON, password
func (_m *CSA) Import(keyJSON []byte, password string) (csakey.KeyV2, error) {
	ret := _m.Called(keyJSON, password)

	var r0 csakey.KeyV2
	if rf, ok := ret.Get(0).(func([]byte, string) csakey.KeyV2); ok {
		r0 = rf(keyJSON, password)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCSA interface {
	mock.TestingT
	Cleanup(func())
}

// NewCSA creates a new instance of CSA. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCSA(t mockConstructorTestingTNewCSA) *CSA {
	mock := &CSA{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
