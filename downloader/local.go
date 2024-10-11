package downloader

import (
	"context"
	"os"

	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/sync"
)

func DownloadLocal(ctx context.Context, src, dst string) error {
	var err error
	fsInfo, err := fs.NewFs(ctx, src)
	if err != nil {
		return err
	}
	if err = os.MkdirAll(dst, os.ModePerm); err != nil {
		return err
	}

	destinationFs, err := fs.NewFs(ctx, dst)
	if err != nil {
		return err
	}

	err = sync.CopyDir(ctx, destinationFs, fsInfo, false)
	if err != nil {
		return err
	}
	return nil
}
