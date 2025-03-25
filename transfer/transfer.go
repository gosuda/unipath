package transfer

import (
	"context"
	"fmt"
	"io"

	"gosuda.org/unipath/transfer/ipfs"
	"gosuda.org/unipath/transfer/rclone"
	"gosuda.org/unipath/transfer/torrent"
	"gosuda.org/unipath/unipath"
)

func Transfer(ctx context.Context, src, dst string, opts ...Option) error {
	var cfg config
	for _, apply := range opts {
		apply(&cfg)
	}

	from, err := unipath.NewUniPathFromString(src)
	if err != nil {
		return fmt.Errorf("failed to parse src uri: %w", err)
	}
	to, err := unipath.NewUniPathFromString(dst)
	if err != nil {
		return fmt.Errorf("failed to parse dst uri: %w", err)
	}
	srcHandler, err := getHandler(from.Protocol)
	if err != nil {
		return err
	}
	dstHandler, err := getHandler(to.Protocol)
	if err != nil {
		return err
	}

	r, err := srcHandler.Read(ctx, from)
	if err != nil {
		return fmt.Errorf("read failed: %w", err)
	}
	defer r.Close()

	reader := io.Reader(r)
	if cfg.onProgress != nil {
		reader = wrapProgressReader(r, cfg.onProgress)
	}

	if err := dstHandler.Write(ctx, to, reader); err != nil {
		return fmt.Errorf("write failed: %w", err)
	}

	if cfg.onComplete != nil {
		cfg.onComplete()
	}
	return nil
}

type ProtocolHandler interface {
	Read(ctx context.Context, from *unipath.UniPath) (io.ReadCloser, error)
	Write(ctx context.Context, to *unipath.UniPath, reader io.Reader) error
}

func getHandler(p unipath.Protocol) (ProtocolHandler, error) {
	switch p {
	case unipath.Local, unipath.Http, unipath.Https, unipath.Ftp, unipath.Sftp,
		unipath.S3, unipath.Dropbox, unipath.GoogleCloudStorage, unipath.Onedrive:
		return &rclone.RcloneHandler{}, nil
	case unipath.Ipfs:
		return &ipfs.IpfsHandler{}, nil
	case unipath.Torrent, unipath.Magnet:
		return &torrent.TorrentHandler{}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", p)
	}
}
