package local

import (
	"errors"

	"gosuda.org/unipath/services"
)

type Config struct {
	BasePath string
}

func New(c Config) (services.Directory, error) {
	rootDir := Directory{
		Object{
			Path: c.BasePath,
		},
	}

	// Checking if we can read from the directory
	if _, err := rootDir.Stat(); err != nil {
		return nil, errors.New("couldn't read, either directory or it's permission is invalid")
	}

	return &rootDir, nil
}
