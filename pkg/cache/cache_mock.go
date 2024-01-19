// Code generated by mockery v2.33.0. DO NOT EDIT.

package cache

import (
	context "context"

	rbac "github.com/RedHatInsights/rbac-client-go"
	mock "github.com/stretchr/testify/mock"
)

// MockCache is an autogenerated mock type for the Cache type
type MockCache struct {
	mock.Mock
}

// GetAccessList provides a mock function with given fields: ctx
func (_m *MockCache) GetAccessList(ctx context.Context) (rbac.AccessList, error) {
	ret := _m.Called(ctx)

	var r0 rbac.AccessList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (rbac.AccessList, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) rbac.AccessList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rbac.AccessList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPulpContentPath provides a mock function with given fields: ctx
func (_m *MockCache) GetPulpContentPath(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAccessList provides a mock function with given fields: ctx, accessList
func (_m *MockCache) SetAccessList(ctx context.Context, accessList rbac.AccessList) error {
	ret := _m.Called(ctx, accessList)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, rbac.AccessList) error); ok {
		r0 = rf(ctx, accessList)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPulpContentPath provides a mock function with given fields: ctx, pulpContentPath
func (_m *MockCache) SetPulpContentPath(ctx context.Context, pulpContentPath string) error {
	ret := _m.Called(ctx, pulpContentPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, pulpContentPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockCache creates a new instance of MockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCache {
	mock := &MockCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
