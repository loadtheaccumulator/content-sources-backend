// Code generated by mockery v2.46.0. DO NOT EDIT.

package dao

import (
	context "context"

	api "github.com/content-services/content-sources-backend/pkg/api"

	mock "github.com/stretchr/testify/mock"
)

// MockAdminTaskDao is an autogenerated mock type for the AdminTaskDao type
type MockAdminTaskDao struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, id
func (_m *MockAdminTaskDao) Fetch(ctx context.Context, id string) (api.AdminTaskInfoResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 api.AdminTaskInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (api.AdminTaskInfoResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) api.AdminTaskInfoResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(api.AdminTaskInfoResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, pageData, filterData
func (_m *MockAdminTaskDao) List(ctx context.Context, pageData api.PaginationData, filterData api.AdminTaskFilterData) (api.AdminTaskInfoCollectionResponse, int64, error) {
	ret := _m.Called(ctx, pageData, filterData)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 api.AdminTaskInfoCollectionResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, api.PaginationData, api.AdminTaskFilterData) (api.AdminTaskInfoCollectionResponse, int64, error)); ok {
		return rf(ctx, pageData, filterData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.PaginationData, api.AdminTaskFilterData) api.AdminTaskInfoCollectionResponse); ok {
		r0 = rf(ctx, pageData, filterData)
	} else {
		r0 = ret.Get(0).(api.AdminTaskInfoCollectionResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.PaginationData, api.AdminTaskFilterData) int64); ok {
		r1 = rf(ctx, pageData, filterData)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, api.PaginationData, api.AdminTaskFilterData) error); ok {
		r2 = rf(ctx, pageData, filterData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewMockAdminTaskDao creates a new instance of MockAdminTaskDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAdminTaskDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAdminTaskDao {
	mock := &MockAdminTaskDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
