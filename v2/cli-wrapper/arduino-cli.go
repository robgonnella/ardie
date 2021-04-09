package cli

import (
	"context"
	"io"

	"github.com/arduino/arduino-cli/cli/globals"
	"github.com/arduino/arduino-cli/cli/instance"
	"github.com/arduino/arduino-cli/commands"
	"github.com/arduino/arduino-cli/commands/board"
	"github.com/arduino/arduino-cli/commands/compile"
	"github.com/arduino/arduino-cli/commands/core"
	"github.com/arduino/arduino-cli/commands/lib"
	"github.com/arduino/arduino-cli/commands/upload"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
)

// Cli reprents our wrapper around arduino-cli
//go:generate mockgen -destination=../mocks/mock_cli.go -package=mocks github.com/robgonnella/ardi/v2/cli-wrapper Cli
type Cli interface {
	CreateInstance() (*rpc.Instance, error)
	CreateInstanceIgnorePlatformIndexErrors() *rpc.Instance
	UpdateIndex(context.Context, *rpc.UpdateIndexRequest, commands.DownloadProgressCB) (*rpc.UpdateIndexResponse, error)
	UpdateLibrariesIndex(context.Context, *rpc.UpdateLibrariesIndexRequest, commands.DownloadProgressCB) error
	PlatformUpgrade(context.Context, *rpc.PlatformUpgradeRequest, commands.DownloadProgressCB, commands.TaskProgressCB) (*rpc.PlatformUpgradeResponse, error)
	PlatformInstall(context.Context, *rpc.PlatformInstallRequest, commands.DownloadProgressCB, commands.TaskProgressCB) (*rpc.PlatformInstallResponse, error)
	PlatformUninstall(context.Context, *rpc.PlatformUninstallRequest, func(curr *rpc.TaskProgress)) (*rpc.PlatformUninstallResponse, error)
	GetPlatforms(*rpc.PlatformListRequest) ([]*rpc.Platform, error)
	PlatformSearch(*rpc.PlatformSearchRequest) (*rpc.PlatformSearchResponse, error)
	ConnectedBoards(instanceID int32) (r []*rpc.DetectedPort, e error)
	Upload(context.Context, *rpc.UploadRequest, io.Writer, io.Writer) (*rpc.UploadResponse, error)
	Compile(context.Context, *rpc.CompileRequest, io.Writer, io.Writer, bool) (*rpc.CompileResponse, error)
	LibrarySearch(context.Context, *rpc.LibrarySearchRequest) (*rpc.LibrarySearchResponse, error)
	LibraryInstall(context.Context, *rpc.LibraryInstallRequest, commands.DownloadProgressCB, commands.TaskProgressCB) error
	LibraryUninstall(context.Context, *rpc.LibraryUninstallRequest, commands.TaskProgressCB) error
	LibraryList(context.Context, *rpc.LibraryListRequest) (*rpc.LibraryListResponse, error)
	Version() string
}

// ArduinoCli represents our wrapper around arduino-cli
type ArduinoCli struct{}

func newArduinoCli() *ArduinoCli {
	return &ArduinoCli{}
}

// CreateInstance wrapper around arduino-cli CreateInstance
func (c *ArduinoCli) CreateInstance() (*rpc.Instance, error) {
	return instance.CreateInstance()
}

// CreateInstanceIgnorePlatformIndexErrors wrapper around arduino-cli CreateInstanceIgnorePlatformIndexErrors
func (c *ArduinoCli) CreateInstanceIgnorePlatformIndexErrors() *rpc.Instance {
	return instance.CreateInstanceIgnorePlatformIndexErrors()
}

// UpdateIndex wrapper around arduino-cli UpdateIndex
func (c *ArduinoCli) UpdateIndex(ctx context.Context, req *rpc.UpdateIndexRequest, fn commands.DownloadProgressCB) (*rpc.UpdateIndexResponse, error) {
	return commands.UpdateIndex(ctx, req, fn)
}

// UpdateLibrariesIndex wrapper around arduino-cli UpdateLibrariesIndex
func (c *ArduinoCli) UpdateLibrariesIndex(ctx context.Context, req *rpc.UpdateLibrariesIndexRequest, fn commands.DownloadProgressCB) error {
	return commands.UpdateLibrariesIndex(ctx, req, fn)
}

// PlatformUpgrade wrapper around arduino-cli PlatformUpgrade
func (c *ArduinoCli) PlatformUpgrade(ctx context.Context, req *rpc.PlatformUpgradeRequest, dlfn commands.DownloadProgressCB, tfn commands.TaskProgressCB) (*rpc.PlatformUpgradeResponse, error) {
	return core.PlatformUpgrade(ctx, req, dlfn, tfn)
}

// PlatformInstall wrapper around arduino-cli PlatformInstall
func (c *ArduinoCli) PlatformInstall(ctx context.Context, req *rpc.PlatformInstallRequest, dlfn commands.DownloadProgressCB, tfn commands.TaskProgressCB) (*rpc.PlatformInstallResponse, error) {
	return core.PlatformInstall(ctx, req, dlfn, tfn)
}

// PlatformUninstall wrapper around arduino-cli PlatformUninstall
func (c *ArduinoCli) PlatformUninstall(ctx context.Context, req *rpc.PlatformUninstallRequest, fn func(curr *rpc.TaskProgress)) (*rpc.PlatformUninstallResponse, error) {
	return core.PlatformUninstall(ctx, req, fn)
}

// GetPlatforms wrapper around arduino-cli GetPlatforms
func (c *ArduinoCli) GetPlatforms(req *rpc.PlatformListRequest) ([]*rpc.Platform, error) {
	return core.GetPlatforms(req)
}

// PlatformSearch wrapper around arduino-cli PlatformSearch
func (c *ArduinoCli) PlatformSearch(req *rpc.PlatformSearchRequest) (*rpc.PlatformSearchResponse, error) {
	return core.PlatformSearch(req)
}

// ConnectedBoards wrapper around arduino-cli board.List
func (c *ArduinoCli) ConnectedBoards(instanceID int32) (r []*rpc.DetectedPort, e error) {
	return board.List(instanceID)
}

// Upload wrapper around arduino-cli Upload
func (c *ArduinoCli) Upload(ctx context.Context, req *rpc.UploadRequest, out io.Writer, err io.Writer) (*rpc.UploadResponse, error) {
	return upload.Upload(ctx, req, out, err)
}

// Compile wrapper around arduino-cli Compile
func (c *ArduinoCli) Compile(ctx context.Context, req *rpc.CompileRequest, out io.Writer, err io.Writer, verbose bool) (*rpc.CompileResponse, error) {
	return compile.Compile(ctx, req, out, err, verbose)
}

// LibrarySearch wrapper around arduino-cli LibrarySearch
func (c *ArduinoCli) LibrarySearch(ctx context.Context, req *rpc.LibrarySearchRequest) (*rpc.LibrarySearchResponse, error) {
	return lib.LibrarySearch(ctx, req)
}

// LibraryInstall wrapper around arduino-cli LibraryInstall
func (c *ArduinoCli) LibraryInstall(ctx context.Context, req *rpc.LibraryInstallRequest, dlfn commands.DownloadProgressCB, tfn commands.TaskProgressCB) error {
	return lib.LibraryInstall(ctx, req, dlfn, tfn)
}

// LibraryUninstall wrapper around arduino-cli LibraryUninstall
func (c *ArduinoCli) LibraryUninstall(ctx context.Context, req *rpc.LibraryUninstallRequest, tfn commands.TaskProgressCB) error {
	return lib.LibraryUninstall(ctx, req, tfn)
}

// LibraryList wrapper around arduino-cli LibraryList
func (c *ArduinoCli) LibraryList(ctx context.Context, req *rpc.LibraryListRequest) (*rpc.LibraryListResponse, error) {
	return lib.LibraryList(ctx, req)
}

// Version wrapper around arduino-cli global version
func (c *ArduinoCli) Version() string {
	return globals.VersionInfo.String()
}
