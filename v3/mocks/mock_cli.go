// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/robgonnella/ardi/v3/cli-wrapper (interfaces: Cli)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	commands "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockCli is a mock of Cli interface.
type MockCli struct {
	ctrl     *gomock.Controller
	recorder *MockCliMockRecorder
}

// MockCliMockRecorder is the mock recorder for MockCli.
type MockCliMockRecorder struct {
	mock *MockCli
}

// NewMockCli creates a new mock instance.
func NewMockCli(ctrl *gomock.Controller) *MockCli {
	mock := &MockCli{ctrl: ctrl}
	mock.recorder = &MockCliMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCli) EXPECT() *MockCliMockRecorder {
	return m.recorder
}

// CreateInstance mocks base method.
func (m *MockCli) CreateInstance() *commands.Instance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance")
	ret0, _ := ret[0].(*commands.Instance)
	return ret0
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockCliMockRecorder) CreateInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockCli)(nil).CreateInstance))
}

// GetPlatforms mocks base method.
func (m *MockCli) GetPlatforms(arg0 *commands.PlatformListRequest) ([]*commands.Platform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatforms", arg0)
	ret0, _ := ret[0].([]*commands.Platform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlatforms indicates an expected call of GetPlatforms.
func (mr *MockCliMockRecorder) GetPlatforms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatforms", reflect.TypeOf((*MockCli)(nil).GetPlatforms), arg0)
}

// InitSettings mocks base method.
func (m *MockCli) InitSettings(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitSettings", arg0)
}

// InitSettings indicates an expected call of InitSettings.
func (mr *MockCliMockRecorder) InitSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitSettings", reflect.TypeOf((*MockCli)(nil).InitSettings), arg0)
}

// LibraryInstall mocks base method.
func (m *MockCli) LibraryInstall(arg0 context.Context, arg1 *commands.LibraryInstallRequest, arg2 commands.DownloadProgressCB, arg3 commands.TaskProgressCB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LibraryInstall", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// LibraryInstall indicates an expected call of LibraryInstall.
func (mr *MockCliMockRecorder) LibraryInstall(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LibraryInstall", reflect.TypeOf((*MockCli)(nil).LibraryInstall), arg0, arg1, arg2, arg3)
}

// LibraryList mocks base method.
func (m *MockCli) LibraryList(arg0 context.Context, arg1 *commands.LibraryListRequest) (*commands.LibraryListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LibraryList", arg0, arg1)
	ret0, _ := ret[0].(*commands.LibraryListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LibraryList indicates an expected call of LibraryList.
func (mr *MockCliMockRecorder) LibraryList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LibraryList", reflect.TypeOf((*MockCli)(nil).LibraryList), arg0, arg1)
}

// LibrarySearch mocks base method.
func (m *MockCli) LibrarySearch(arg0 context.Context, arg1 *commands.LibrarySearchRequest) (*commands.LibrarySearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LibrarySearch", arg0, arg1)
	ret0, _ := ret[0].(*commands.LibrarySearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LibrarySearch indicates an expected call of LibrarySearch.
func (mr *MockCliMockRecorder) LibrarySearch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LibrarySearch", reflect.TypeOf((*MockCli)(nil).LibrarySearch), arg0, arg1)
}

// LibraryUninstall mocks base method.
func (m *MockCli) LibraryUninstall(arg0 context.Context, arg1 *commands.LibraryUninstallRequest, arg2 commands.TaskProgressCB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LibraryUninstall", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// LibraryUninstall indicates an expected call of LibraryUninstall.
func (mr *MockCliMockRecorder) LibraryUninstall(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LibraryUninstall", reflect.TypeOf((*MockCli)(nil).LibraryUninstall), arg0, arg1, arg2)
}

// PlatformInstall mocks base method.
func (m *MockCli) PlatformInstall(arg0 context.Context, arg1 *commands.PlatformInstallRequest, arg2 commands.DownloadProgressCB, arg3 commands.TaskProgressCB) (*commands.PlatformInstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlatformInstall", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*commands.PlatformInstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlatformInstall indicates an expected call of PlatformInstall.
func (mr *MockCliMockRecorder) PlatformInstall(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatformInstall", reflect.TypeOf((*MockCli)(nil).PlatformInstall), arg0, arg1, arg2, arg3)
}

// PlatformSearch mocks base method.
func (m *MockCli) PlatformSearch(arg0 *commands.PlatformSearchRequest) (*commands.PlatformSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlatformSearch", arg0)
	ret0, _ := ret[0].(*commands.PlatformSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlatformSearch indicates an expected call of PlatformSearch.
func (mr *MockCliMockRecorder) PlatformSearch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatformSearch", reflect.TypeOf((*MockCli)(nil).PlatformSearch), arg0)
}

// PlatformUninstall mocks base method.
func (m *MockCli) PlatformUninstall(arg0 context.Context, arg1 *commands.PlatformUninstallRequest, arg2 func(*commands.TaskProgress)) (*commands.PlatformUninstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlatformUninstall", arg0, arg1, arg2)
	ret0, _ := ret[0].(*commands.PlatformUninstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlatformUninstall indicates an expected call of PlatformUninstall.
func (mr *MockCliMockRecorder) PlatformUninstall(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatformUninstall", reflect.TypeOf((*MockCli)(nil).PlatformUninstall), arg0, arg1, arg2)
}

// UpdateIndex mocks base method.
func (m *MockCli) UpdateIndex(arg0 context.Context, arg1 *commands.UpdateIndexRequest, arg2 commands.DownloadProgressCB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIndex indicates an expected call of UpdateIndex.
func (mr *MockCliMockRecorder) UpdateIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIndex", reflect.TypeOf((*MockCli)(nil).UpdateIndex), arg0, arg1, arg2)
}

// UpdateLibrariesIndex mocks base method.
func (m *MockCli) UpdateLibrariesIndex(arg0 context.Context, arg1 *commands.UpdateLibrariesIndexRequest, arg2 commands.DownloadProgressCB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLibrariesIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLibrariesIndex indicates an expected call of UpdateLibrariesIndex.
func (mr *MockCliMockRecorder) UpdateLibrariesIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLibrariesIndex", reflect.TypeOf((*MockCli)(nil).UpdateLibrariesIndex), arg0, arg1, arg2)
}

// Version mocks base method.
func (m *MockCli) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockCliMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockCli)(nil).Version))
}
