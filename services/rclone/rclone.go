package rclone

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/cache"
	"github.com/rclone/rclone/fs/fspath"
	"github.com/rclone/rclone/fs/operations"
	"github.com/rclone/rclone/fs/sync"
	"gosuda.org/unipath/unipath"
)

type RcloneHandler struct {
}

func (r *RcloneHandler) Read(ctx context.Context, from *unipath.UniPath) (io.ReadCloser, error) {
	return nil, nil
}

func (r *RcloneHandler) Write(ctx context.Context, to *unipath.UniPath, reader io.Reader) error {
	return nil
}

var (
	ErrCopyDirectoryToFile = errors.New("can't copy a directory to a file")
)

func rcloneDownloadLocal(ctx context.Context, src, dst string) error {
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

func rcloneDownloadUrl(ctx context.Context, src, dst string) error {
	var err error
	fdst, dstFileName := NewFsFile(dst)
	if dstFileName != "" {
		_, err = operations.CopyURL(ctx, fdst, dstFileName, src, false, false, false)
	} else {
		_, err = operations.CopyURL(ctx, fdst, dstFileName, src, true, true, false)
	}
	return err
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
