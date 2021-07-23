// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/robgonnella/ardi/v2/cli-wrapper (interfaces: Cli)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	commands "github.com/arduino/arduino-cli/commands"
	commands0 "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
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

// Compile mocks base method.
func (m *MockCli) Compile(arg0 context.Context, arg1 *commands0.CompileRequest, arg2, arg3 io.Writer, arg4 bool) (*commands0.CompileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compile", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*commands0.CompileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compile indicates an expected call of Compile.
func (mr *MockCliMockRecorder) Compile(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compile", reflect.TypeOf((*MockCli)(nil).Compile), arg0, arg1, arg2, arg3, arg4)
}

// ConnectedBoards mocks base method.
func (m *MockCli) ConnectedBoards(arg0 int32) ([]*commands0.DetectedPort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedBoards", arg0)
	ret0, _ := ret[0].([]*commands0.DetectedPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectedBoards indicates an expected call of ConnectedBoards.
func (mr *MockCliMockRecorder) ConnectedBoards(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedBoards", reflect.TypeOf((*MockCli)(nil).ConnectedBoards), arg0)
}

// CreateInstance mocks base method.
func (m *MockCli) CreateInstance() *commands0.Instance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance")
	ret0, _ := ret[0].(*commands0.Instance)
	return ret0
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockCliMockRecorder) CreateInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockCli)(nil).CreateInstance))
}

// GetPlatforms mocks base method.
func (m *MockCli) GetPlatforms(arg0 *commands0.PlatformListRequest) ([]*commands0.Platform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatforms", arg0)
	ret0, _ := ret[0].([]*commands0.Platform)
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
func (m *MockCli) LibraryInstall(arg0 context.Context, arg1 *commands0.LibraryInstallRequest, arg2 commands.DownloadProgressCB, arg3 commands.TaskProgressCB) error {
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
func (m *MockCli) LibraryList(arg0 context.Context, arg1 *commands0.LibraryListRequest) (*commands0.LibraryListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LibraryList", arg0, arg1)
	ret0, _ := ret[0].(*commands0.LibraryListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LibraryList indicates an expected call of LibraryList.
func (mr *MockCliMockRecorder) LibraryList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LibraryList", reflect.TypeOf((*MockCli)(nil).LibraryList), arg0, arg1)
}

// LibrarySearch mocks base method.
func (m *MockCli) LibrarySearch(arg0 context.Context, arg1 *commands0.LibrarySearchRequest) (*commands0.LibrarySearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LibrarySearch", arg0, arg1)
	ret0, _ := ret[0].(*commands0.LibrarySearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LibrarySearch indicates an expected call of LibrarySearch.
func (mr *MockCliMockRecorder) LibrarySearch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LibrarySearch", reflect.TypeOf((*MockCli)(nil).LibrarySearch), arg0, arg1)
}

// LibraryUninstall mocks base method.
func (m *MockCli) LibraryUninstall(arg0 context.Context, arg1 *commands0.LibraryUninstallRequest, arg2 commands.TaskProgressCB) error {
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
func (m *MockCli) PlatformInstall(arg0 context.Context, arg1 *commands0.PlatformInstallRequest, arg2 commands.DownloadProgressCB, arg3 commands.TaskProgressCB) (*commands0.PlatformInstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlatformInstall", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*commands0.PlatformInstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlatformInstall indicates an expected call of PlatformInstall.
func (mr *MockCliMockRecorder) PlatformInstall(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatformInstall", reflect.TypeOf((*MockCli)(nil).PlatformInstall), arg0, arg1, arg2, arg3)
}

// PlatformSearch mocks base method.
func (m *MockCli) PlatformSearch(arg0 *commands0.PlatformSearchRequest) (*commands0.PlatformSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlatformSearch", arg0)
	ret0, _ := ret[0].(*commands0.PlatformSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlatformSearch indicates an expected call of PlatformSearch.
func (mr *MockCliMockRecorder) PlatformSearch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatformSearch", reflect.TypeOf((*MockCli)(nil).PlatformSearch), arg0)
}

// PlatformUninstall mocks base method.
func (m *MockCli) PlatformUninstall(arg0 context.Context, arg1 *commands0.PlatformUninstallRequest, arg2 func(*commands0.TaskProgress)) (*commands0.PlatformUninstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlatformUninstall", arg0, arg1, arg2)
	ret0, _ := ret[0].(*commands0.PlatformUninstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlatformUninstall indicates an expected call of PlatformUninstall.
func (mr *MockCliMockRecorder) PlatformUninstall(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatformUninstall", reflect.TypeOf((*MockCli)(nil).PlatformUninstall), arg0, arg1, arg2)
}

// UpdateIndex mocks base method.
func (m *MockCli) UpdateIndex(arg0 context.Context, arg1 *commands0.UpdateIndexRequest, arg2 commands.DownloadProgressCB) (*commands0.UpdateIndexResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(*commands0.UpdateIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIndex indicates an expected call of UpdateIndex.
func (mr *MockCliMockRecorder) UpdateIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIndex", reflect.TypeOf((*MockCli)(nil).UpdateIndex), arg0, arg1, arg2)
}

// UpdateLibrariesIndex mocks base method.
func (m *MockCli) UpdateLibrariesIndex(arg0 context.Context, arg1 *commands0.UpdateLibrariesIndexRequest, arg2 commands.DownloadProgressCB) error {
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

// Upload mocks base method.
func (m *MockCli) Upload(arg0 context.Context, arg1 *commands0.UploadRequest, arg2, arg3 io.Writer) (*commands0.UploadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*commands0.UploadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockCliMockRecorder) Upload(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockCli)(nil).Upload), arg0, arg1, arg2, arg3)
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
