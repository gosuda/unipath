package torrent

import (
	"context"
	"io"

	"gosuda.org/unipath/unipath"
)

type TorrentHandler struct {
}

func (r *TorrentHandler) Read(ctx context.Context, from *unipath.UniPath) (io.ReadCloser, error) {
	return nil, nil
}

func (r *TorrentHandler) Write(ctx context.Context, to *unipath.UniPath, reader io.Reader) error {
	return nil
}
