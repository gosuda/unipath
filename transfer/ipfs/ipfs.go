package ipfs

import (
	"context"
	"io"

	"gosuda.org/unipath/unipath"
)

type IpfsHandler struct {
}

func (r *IpfsHandler) Read(ctx context.Context, from *unipath.UniPath) (io.ReadCloser, error) {
	return nil, nil
}

func (r *IpfsHandler) Write(ctx context.Context, to *unipath.UniPath, reader io.Reader) error {
	return nil
}
