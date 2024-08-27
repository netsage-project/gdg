// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	filters "github.com/esnet/gdg/internal/service/filters"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana-openapi-client-go/models"

	types "github.com/esnet/gdg/internal/types"
)

// ConnectionsApi is an autogenerated mock type for the ConnectionsApi type
type ConnectionsApi struct {
	mock.Mock
}

type ConnectionsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *ConnectionsApi) EXPECT() *ConnectionsApi_Expecter {
	return &ConnectionsApi_Expecter{mock: &_m.Mock}
}

// DeleteAllConnectionPermissions provides a mock function with given fields: filter
func (_m *ConnectionsApi) DeleteAllConnectionPermissions(filter filters.Filter) []string {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllConnectionPermissions")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ConnectionsApi_DeleteAllConnectionPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAllConnectionPermissions'
type ConnectionsApi_DeleteAllConnectionPermissions_Call struct {
	*mock.Call
}

// DeleteAllConnectionPermissions is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) DeleteAllConnectionPermissions(filter interface{}) *ConnectionsApi_DeleteAllConnectionPermissions_Call {
	return &ConnectionsApi_DeleteAllConnectionPermissions_Call{Call: _e.mock.On("DeleteAllConnectionPermissions", filter)}
}

func (_c *ConnectionsApi_DeleteAllConnectionPermissions_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_DeleteAllConnectionPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_DeleteAllConnectionPermissions_Call) Return(_a0 []string) *ConnectionsApi_DeleteAllConnectionPermissions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_DeleteAllConnectionPermissions_Call) RunAndReturn(run func(filters.Filter) []string) *ConnectionsApi_DeleteAllConnectionPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAllConnections provides a mock function with given fields: filter
func (_m *ConnectionsApi) DeleteAllConnections(filter filters.Filter) []string {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllConnections")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ConnectionsApi_DeleteAllConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAllConnections'
type ConnectionsApi_DeleteAllConnections_Call struct {
	*mock.Call
}

// DeleteAllConnections is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) DeleteAllConnections(filter interface{}) *ConnectionsApi_DeleteAllConnections_Call {
	return &ConnectionsApi_DeleteAllConnections_Call{Call: _e.mock.On("DeleteAllConnections", filter)}
}

func (_c *ConnectionsApi_DeleteAllConnections_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_DeleteAllConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_DeleteAllConnections_Call) Return(_a0 []string) *ConnectionsApi_DeleteAllConnections_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_DeleteAllConnections_Call) RunAndReturn(run func(filters.Filter) []string) *ConnectionsApi_DeleteAllConnections_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadConnectionPermissions provides a mock function with given fields: filter
func (_m *ConnectionsApi) DownloadConnectionPermissions(filter filters.Filter) []string {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for DownloadConnectionPermissions")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ConnectionsApi_DownloadConnectionPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadConnectionPermissions'
type ConnectionsApi_DownloadConnectionPermissions_Call struct {
	*mock.Call
}

// DownloadConnectionPermissions is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) DownloadConnectionPermissions(filter interface{}) *ConnectionsApi_DownloadConnectionPermissions_Call {
	return &ConnectionsApi_DownloadConnectionPermissions_Call{Call: _e.mock.On("DownloadConnectionPermissions", filter)}
}

func (_c *ConnectionsApi_DownloadConnectionPermissions_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_DownloadConnectionPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_DownloadConnectionPermissions_Call) Return(_a0 []string) *ConnectionsApi_DownloadConnectionPermissions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_DownloadConnectionPermissions_Call) RunAndReturn(run func(filters.Filter) []string) *ConnectionsApi_DownloadConnectionPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadConnections provides a mock function with given fields: filter
func (_m *ConnectionsApi) DownloadConnections(filter filters.Filter) []string {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for DownloadConnections")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ConnectionsApi_DownloadConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadConnections'
type ConnectionsApi_DownloadConnections_Call struct {
	*mock.Call
}

// DownloadConnections is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) DownloadConnections(filter interface{}) *ConnectionsApi_DownloadConnections_Call {
	return &ConnectionsApi_DownloadConnections_Call{Call: _e.mock.On("DownloadConnections", filter)}
}

func (_c *ConnectionsApi_DownloadConnections_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_DownloadConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_DownloadConnections_Call) Return(_a0 []string) *ConnectionsApi_DownloadConnections_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_DownloadConnections_Call) RunAndReturn(run func(filters.Filter) []string) *ConnectionsApi_DownloadConnections_Call {
	_c.Call.Return(run)
	return _c
}

// ListConnectionPermissions provides a mock function with given fields: filter
func (_m *ConnectionsApi) ListConnectionPermissions(filter filters.Filter) []types.ConnectionPermissionItem {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for ListConnectionPermissions")
	}

	var r0 []types.ConnectionPermissionItem
	if rf, ok := ret.Get(0).(func(filters.Filter) []types.ConnectionPermissionItem); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ConnectionPermissionItem)
		}
	}

	return r0
}

// ConnectionsApi_ListConnectionPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListConnectionPermissions'
type ConnectionsApi_ListConnectionPermissions_Call struct {
	*mock.Call
}

// ListConnectionPermissions is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) ListConnectionPermissions(filter interface{}) *ConnectionsApi_ListConnectionPermissions_Call {
	return &ConnectionsApi_ListConnectionPermissions_Call{Call: _e.mock.On("ListConnectionPermissions", filter)}
}

func (_c *ConnectionsApi_ListConnectionPermissions_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_ListConnectionPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_ListConnectionPermissions_Call) Return(_a0 []types.ConnectionPermissionItem) *ConnectionsApi_ListConnectionPermissions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_ListConnectionPermissions_Call) RunAndReturn(run func(filters.Filter) []types.ConnectionPermissionItem) *ConnectionsApi_ListConnectionPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// ListConnections provides a mock function with given fields: filter
func (_m *ConnectionsApi) ListConnections(filter filters.Filter) []models.DataSourceListItemDTO {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for ListConnections")
	}

	var r0 []models.DataSourceListItemDTO
	if rf, ok := ret.Get(0).(func(filters.Filter) []models.DataSourceListItemDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DataSourceListItemDTO)
		}
	}

	return r0
}

// ConnectionsApi_ListConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListConnections'
type ConnectionsApi_ListConnections_Call struct {
	*mock.Call
}

// ListConnections is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) ListConnections(filter interface{}) *ConnectionsApi_ListConnections_Call {
	return &ConnectionsApi_ListConnections_Call{Call: _e.mock.On("ListConnections", filter)}
}

func (_c *ConnectionsApi_ListConnections_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_ListConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_ListConnections_Call) Return(_a0 []models.DataSourceListItemDTO) *ConnectionsApi_ListConnections_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_ListConnections_Call) RunAndReturn(run func(filters.Filter) []models.DataSourceListItemDTO) *ConnectionsApi_ListConnections_Call {
	_c.Call.Return(run)
	return _c
}

// UploadConnectionPermissions provides a mock function with given fields: filter
func (_m *ConnectionsApi) UploadConnectionPermissions(filter filters.Filter) []string {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for UploadConnectionPermissions")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ConnectionsApi_UploadConnectionPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadConnectionPermissions'
type ConnectionsApi_UploadConnectionPermissions_Call struct {
	*mock.Call
}

// UploadConnectionPermissions is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) UploadConnectionPermissions(filter interface{}) *ConnectionsApi_UploadConnectionPermissions_Call {
	return &ConnectionsApi_UploadConnectionPermissions_Call{Call: _e.mock.On("UploadConnectionPermissions", filter)}
}

func (_c *ConnectionsApi_UploadConnectionPermissions_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_UploadConnectionPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_UploadConnectionPermissions_Call) Return(_a0 []string) *ConnectionsApi_UploadConnectionPermissions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_UploadConnectionPermissions_Call) RunAndReturn(run func(filters.Filter) []string) *ConnectionsApi_UploadConnectionPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// UploadConnections provides a mock function with given fields: filter
func (_m *ConnectionsApi) UploadConnections(filter filters.Filter) []string {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for UploadConnections")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ConnectionsApi_UploadConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadConnections'
type ConnectionsApi_UploadConnections_Call struct {
	*mock.Call
}

// UploadConnections is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *ConnectionsApi_Expecter) UploadConnections(filter interface{}) *ConnectionsApi_UploadConnections_Call {
	return &ConnectionsApi_UploadConnections_Call{Call: _e.mock.On("UploadConnections", filter)}
}

func (_c *ConnectionsApi_UploadConnections_Call) Run(run func(filter filters.Filter)) *ConnectionsApi_UploadConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *ConnectionsApi_UploadConnections_Call) Return(_a0 []string) *ConnectionsApi_UploadConnections_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectionsApi_UploadConnections_Call) RunAndReturn(run func(filters.Filter) []string) *ConnectionsApi_UploadConnections_Call {
	_c.Call.Return(run)
	return _c
}

// NewConnectionsApi creates a new instance of ConnectionsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnectionsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConnectionsApi {
	mock := &ConnectionsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
