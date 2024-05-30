// Code generated by mockery v2.43.0. DO NOT EDIT.

package dao

import (
	context "context"

	api "github.com/content-services/content-sources-backend/pkg/api"

	mock "github.com/stretchr/testify/mock"
)

// MockTaskInfoDao is an autogenerated mock type for the TaskInfoDao type
type MockTaskInfoDao struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: ctx
func (_m *MockTaskInfoDao) Cleanup(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, OrgID, id
func (_m *MockTaskInfoDao) Fetch(ctx context.Context, OrgID string, id string) (api.TaskInfoResponse, error) {
	ret := _m.Called(ctx, OrgID, id)

	var r0 api.TaskInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (api.TaskInfoResponse, error)); ok {
		return rf(ctx, OrgID, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) api.TaskInfoResponse); ok {
		r0 = rf(ctx, OrgID, id)
	} else {
		r0 = ret.Get(0).(api.TaskInfoResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, OrgID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTaskInProgress provides a mock function with given fields: ctx, orgID, repoUUID, taskType
func (_m *MockTaskInfoDao) IsTaskInProgress(ctx context.Context, orgID string, repoUUID string, taskType string) (bool, string, error) {
	ret := _m.Called(ctx, orgID, repoUUID, taskType)

	var r0 bool
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (bool, string, error)); ok {
		return rf(ctx, orgID, repoUUID, taskType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) bool); ok {
		r0 = rf(ctx, orgID, repoUUID, taskType)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) string); ok {
		r1 = rf(ctx, orgID, repoUUID, taskType)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, string) error); ok {
		r2 = rf(ctx, orgID, repoUUID, taskType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// List provides a mock function with given fields: ctx, OrgID, pageData, filterData
func (_m *MockTaskInfoDao) List(ctx context.Context, OrgID string, pageData api.PaginationData, filterData api.TaskInfoFilterData) (api.TaskInfoCollectionResponse, int64, error) {
	ret := _m.Called(ctx, OrgID, pageData, filterData)

	var r0 api.TaskInfoCollectionResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, api.PaginationData, api.TaskInfoFilterData) (api.TaskInfoCollectionResponse, int64, error)); ok {
		return rf(ctx, OrgID, pageData, filterData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, api.PaginationData, api.TaskInfoFilterData) api.TaskInfoCollectionResponse); ok {
		r0 = rf(ctx, OrgID, pageData, filterData)
	} else {
		r0 = ret.Get(0).(api.TaskInfoCollectionResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, api.PaginationData, api.TaskInfoFilterData) int64); ok {
		r1 = rf(ctx, OrgID, pageData, filterData)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, api.PaginationData, api.TaskInfoFilterData) error); ok {
		r2 = rf(ctx, OrgID, pageData, filterData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewMockTaskInfoDao creates a new instance of MockTaskInfoDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTaskInfoDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTaskInfoDao {
	mock := &MockTaskInfoDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
