package downloader

import (
	"context"
	"errors"
	"os/exec"
	"path"
	"runtime"

	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/cache"
	"github.com/rclone/rclone/fs/fspath"
	"github.com/rclone/rclone/fs/operations"
	"github.com/rclone/rclone/fs/sync"
)

var (
	ErrCopyDirectoryToFile = errors.New("can't copy a directory to a file")
)

func DownloadLocal(ctx context.Context, src, dst string) error {
	fsrc, srcFileName, fdst, dstFileName := NewFsSrcFileDst(src, dst)
	// folder copy
	if srcFileName == "" {
		if dstFileName != "" {
			return ErrCopyDirectoryToFile
		}
		return sync.CopyDir(ctx, fdst, fsrc, false)
	}

	// file copy
	if dstFileName == "" {
		dstFileName = srcFileName
	}
	return operations.CopyFile(ctx, fdst, fsrc, dstFileName, srcFileName)
}

func DownloadUrl(ctx context.Context, src, dst string) error {
	var err error
	fdst, dstFileName := NewFsFile(dst)
	if dstFileName != "" {
		_, err = operations.CopyURL(ctx, fdst, dstFileName, src, false, false, false)
	} else {
		_, err = operations.CopyURL(ctx, fdst, dstFileName, src, true, true, false)
	}
	return err
}

func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}

func NewFsSrcFileDst(src, dst string) (fsrc fs.Fs, srcFileName string, fdst fs.Fs, dstFileName string) {
	fsrc, srcFileName = NewFsFile(src)
	fdst, dstFileName = NewFsFile(dst)
	return fsrc, srcFileName, fdst, dstFileName
}

func NewFsFile(remote string) (fs.Fs, string) {
	ctx := context.Background()
	_, fsPath, err := fspath.SplitFs(remote)
	if err != nil {
		err = fs.CountError(err)
		fs.Fatalf(nil, "Failed to create file system for %q: %v", remote, err)
	}
	f, err := cache.Get(ctx, remote)
	switch err {
	case fs.ErrorIsFile:
		cache.Pin(f) // pin indefinitely since it was on the CLI
		return f, path.Base(fsPath)
	case nil:
		cache.Pin(f) // pin indefinitely since it was on the CLI
		return f, ""
	default:
		err = fs.CountError(err)
		fs.Fatalf(nil, "Failed to create file system for %q: %v", remote, err)
	}
	return nil, ""
}
