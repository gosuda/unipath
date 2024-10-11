package downloader

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func DownloadLocal(ctx context.Context, srcPath, dstPath string) error {
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("could not access source: %w", err)
	}

	if srcInfo.IsDir() {
		return copyDir(ctx, srcPath, dstPath)
	} else {
		return copyFile(ctx, srcPath, dstPath)
	}
}

func copyFile(ctx context.Context, srcPath, dstPath string) error {
	dstDir := filepath.Dir(dstPath)
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = copy(ctx, dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("file copy error: %w", err)
	}

	return nil
}
func copyDir(ctx context.Context, srcDir, dstDir string) error {
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create destination directory: %w", err)
	}

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(ctx, path, dstPath)
	})

	return err
}

func copy(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		return io.Copy(dst, src)
	}
}
