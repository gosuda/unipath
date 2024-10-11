package downloader

import (
	"context"
	"os"
	"path/filepath"

	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/operations"
	"github.com/rclone/rclone/fs/sync"
)

func DownloadLocal(ctx context.Context, src, dst string) error {
	if isDir(src) {
		return copyDir(ctx, src, dst)
	}
	return copyFile(ctx, src, dst)
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

// copyDir: 로컬 디렉토리 복사 전용 함수
func copyDir(ctx context.Context, src, dst string) error {
	srcFs, err := fs.NewFs(ctx, src)
	if err != nil {
		return err
	}

	dstFs, err := fs.NewFs(ctx, dst)
	if err != nil {
		return err
	}

	err = sync.CopyDir(ctx, dstFs, srcFs, true)
	if err != nil {
		return err
	}

	return nil
}

func copyFile(ctx context.Context, src, dst string) error {
	srcFs, err := fs.NewFs(ctx, filepath.Dir(src))
	if err != nil {
		return err
	}

	dstFs, err := fs.NewFs(ctx, filepath.Dir(dst))
	if err != nil {
		return err
	}

	err = operations.CopyFile(ctx, dstFs, srcFs, filepath.Base(dst), filepath.Base(src))
	if err != nil {
		return err
	}

	return nil
}
