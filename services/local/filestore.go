package local

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"sync"

	"gosuda.org/unipath/unipath"
)

func NewClient(rootPath string) *Client {
	return &Client{
		root: Directory{
			Object{
				Path: rootPath,
			},
		},
	}
}

type Client struct {
	root  Directory
	locks sync.Map
}

func (c *Client) getLock(path string) *sync.Mutex {
	lock, _ := c.locks.LoadOrStore(path, &sync.Mutex{})
	return lock.(*sync.Mutex)
}

func (c *Client) Read(ctx context.Context, from *unipath.UniPath) (io.ReadCloser, error) {
	if !from.IsLocal() {
		return nil, fmt.Errorf("unsupported protocol: %s", from.Protocol)
	}

	return nil, nil
	// return file, nil
}

func (c *Client) Write(ctx context.Context, to *unipath.UniPath, reader io.Reader) error {
	if !to.IsLocal() {
		return fmt.Errorf("unsupported protocol: %s", to.Protocol)
	}
	filePath := filepath.Join(c.root.Path, to.Path)

	lock := c.getLock(filePath)
	lock.Lock()
	defer lock.Unlock()

	return nil
}
