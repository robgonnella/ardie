// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/robgonnella/ardi/v2/rpc (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	commands "github.com/arduino/arduino-cli/rpc/commands"
	gomock "github.com/golang/mock/gomock"
	rpc "github.com/robgonnella/ardi/v2/rpc"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AllBoards mocks base method
func (m *MockClient) AllBoards() []*rpc.Board {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllBoards")
	ret0, _ := ret[0].([]*rpc.Board)
	return ret0
}

// AllBoards indicates an expected call of AllBoards
func (mr *MockClientMockRecorder) AllBoards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllBoards", reflect.TypeOf((*MockClient)(nil).AllBoards))
}

// ClientVersion mocks base method
func (m *MockClient) ClientVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientVersion indicates an expected call of ClientVersion
func (mr *MockClientMockRecorder) ClientVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientVersion", reflect.TypeOf((*MockClient)(nil).ClientVersion))
}

// Close mocks base method
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Compile mocks base method
func (m *MockClient) Compile(arg0 rpc.CompileOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Compile indicates an expected call of Compile
func (mr *MockClientMockRecorder) Compile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compile", reflect.TypeOf((*MockClient)(nil).Compile), arg0)
}

// Connect mocks base method
func (m *MockClient) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockClientMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClient)(nil).Connect))
}

// ConnectedBoards mocks base method
func (m *MockClient) ConnectedBoards() []*rpc.Board {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedBoards")
	ret0, _ := ret[0].([]*rpc.Board)
	return ret0
}

// ConnectedBoards indicates an expected call of ConnectedBoards
func (mr *MockClientMockRecorder) ConnectedBoards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedBoards", reflect.TypeOf((*MockClient)(nil).ConnectedBoards))
}

// GetInstalledLibs mocks base method
func (m *MockClient) GetInstalledLibs() ([]*commands.InstalledLibrary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstalledLibs")
	ret0, _ := ret[0].([]*commands.InstalledLibrary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstalledLibs indicates an expected call of GetInstalledLibs
func (mr *MockClientMockRecorder) GetInstalledLibs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstalledLibs", reflect.TypeOf((*MockClient)(nil).GetInstalledLibs))
}

// GetInstalledPlatforms mocks base method
func (m *MockClient) GetInstalledPlatforms() ([]*commands.Platform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstalledPlatforms")
	ret0, _ := ret[0].([]*commands.Platform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstalledPlatforms indicates an expected call of GetInstalledPlatforms
func (mr *MockClientMockRecorder) GetInstalledPlatforms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstalledPlatforms", reflect.TypeOf((*MockClient)(nil).GetInstalledPlatforms))
}

// GetPlatforms mocks base method
func (m *MockClient) GetPlatforms() ([]*commands.Platform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatforms")
	ret0, _ := ret[0].([]*commands.Platform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlatforms indicates an expected call of GetPlatforms
func (mr *MockClientMockRecorder) GetPlatforms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatforms", reflect.TypeOf((*MockClient)(nil).GetPlatforms))
}

// InstallLibrary mocks base method
func (m *MockClient) InstallLibrary(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallLibrary", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallLibrary indicates an expected call of InstallLibrary
func (mr *MockClientMockRecorder) InstallLibrary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallLibrary", reflect.TypeOf((*MockClient)(nil).InstallLibrary), arg0, arg1)
}

// InstallPlatform mocks base method
func (m *MockClient) InstallPlatform(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallPlatform", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InstallPlatform indicates an expected call of InstallPlatform
func (mr *MockClientMockRecorder) InstallPlatform(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallPlatform", reflect.TypeOf((*MockClient)(nil).InstallPlatform), arg0)
}

// SearchLibraries mocks base method
func (m *MockClient) SearchLibraries(arg0 string) ([]*commands.SearchedLibrary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchLibraries", arg0)
	ret0, _ := ret[0].([]*commands.SearchedLibrary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLibraries indicates an expected call of SearchLibraries
func (mr *MockClientMockRecorder) SearchLibraries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLibraries", reflect.TypeOf((*MockClient)(nil).SearchLibraries), arg0)
}

// UninstallLibrary mocks base method
func (m *MockClient) UninstallLibrary(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallLibrary", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallLibrary indicates an expected call of UninstallLibrary
func (mr *MockClientMockRecorder) UninstallLibrary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallLibrary", reflect.TypeOf((*MockClient)(nil).UninstallLibrary), arg0)
}

// UninstallPlatform mocks base method
func (m *MockClient) UninstallPlatform(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallPlatform", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UninstallPlatform indicates an expected call of UninstallPlatform
func (mr *MockClientMockRecorder) UninstallPlatform(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallPlatform", reflect.TypeOf((*MockClient)(nil).UninstallPlatform), arg0)
}

// UpdateIndexFiles mocks base method
func (m *MockClient) UpdateIndexFiles() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIndexFiles")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIndexFiles indicates an expected call of UpdateIndexFiles
func (mr *MockClientMockRecorder) UpdateIndexFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIndexFiles", reflect.TypeOf((*MockClient)(nil).UpdateIndexFiles))
}

// UpdateLibraryIndex mocks base method
func (m *MockClient) UpdateLibraryIndex() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLibraryIndex")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLibraryIndex indicates an expected call of UpdateLibraryIndex
func (mr *MockClientMockRecorder) UpdateLibraryIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLibraryIndex", reflect.TypeOf((*MockClient)(nil).UpdateLibraryIndex))
}

// UpdatePlatformIndex mocks base method
func (m *MockClient) UpdatePlatformIndex() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlatformIndex")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlatformIndex indicates an expected call of UpdatePlatformIndex
func (mr *MockClientMockRecorder) UpdatePlatformIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlatformIndex", reflect.TypeOf((*MockClient)(nil).UpdatePlatformIndex))
}

// UpgradePlatform mocks base method
func (m *MockClient) UpgradePlatform(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradePlatform", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradePlatform indicates an expected call of UpgradePlatform
func (mr *MockClientMockRecorder) UpgradePlatform(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradePlatform", reflect.TypeOf((*MockClient)(nil).UpgradePlatform), arg0)
}

// Upload mocks base method
func (m *MockClient) Upload(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockClientMockRecorder) Upload(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockClient)(nil).Upload), arg0, arg1, arg2)
}
