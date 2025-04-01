package file

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"gosuda.org/unipath/unipath"
)

func NewClient(rootPath string) *Client {
	return &Client{
		rootPath: rootPath,
	}
}

type Client struct {
	rootPath string
	mtx      sync.Mutex
}

func (c *Client) Read(ctx context.Context, from *unipath.UniPath) (io.ReadCloser, error) {
	if from.IsLocal() {
		filePath := filepath.Join(c.rootPath, from.Path)

		c.mtx.Lock()
		defer c.mtx.Unlock()

		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}
		return file, nil
	}
	return nil, fmt.Errorf("unsupported protocol: %s", from.Protocol)
}

func (c *Client) Write(ctx context.Context, to *unipath.UniPath, reader io.Reader) error {
	if to.IsLocal() {
		filePath := filepath.Join(c.rootPath, to.Path)

		c.mtx.Lock()
		defer c.mtx.Unlock()

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()

		_, err = io.Copy(file, reader)
		if err != nil {
			return fmt.Errorf("failed to write file: %v", err)
		}
		return nil
	}
	return fmt.Errorf("unsupported protocol: %s", to.Protocol)
}
