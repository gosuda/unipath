package transfer

import "io"

type Option func(*Config)
type Config struct {
	onProgress func(written int64)
	onComplete func()
}

func WithProgress(fn func(written int64)) Option {
	return func(o *Config) {
		o.onProgress = fn
	}
}

func WithComplete(fn func()) Option {
	return func(o *Config) {
		o.onComplete = fn
	}
}

func wrapProgressReader(r io.Reader, onProgress func(written int64)) io.Reader {
	var total int64 = 0
	buf := make([]byte, 32*1024)

	return &progressReader{
		reader:     r,
		onProgress: onProgress,
		buf:        buf,
		total:      &total,
	}
}

type progressReader struct {
	reader     io.Reader
	onProgress func(written int64)
	buf        []byte
	total      *int64
}

func (p *progressReader) Read(b []byte) (int, error) {
	n, err := p.reader.Read(b)
	if n > 0 {
		*p.total += int64(n)
		p.onProgress(*p.total)
	}
	return n, err
}
