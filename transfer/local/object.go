package local

import (
	"os"

	"gosuda.org/unipath/transfer/iface"
)

type Object struct {
	Path string
}

func (f *Object) GetPath() string {
	return f.Path
}

func (f *Object) String() string {
	return f.GetPath()
}

func (f *Object) Exists() bool {
	if _, err := os.Stat(f.Path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f *Object) Delete() error {
	return os.Remove(f.Path)
}

func (f *Object) Stat() (iface.FileInfo, error) {
	fi, err := os.Stat(f.Path)
	if err != nil {
		return iface.FileInfo{}, err
	}

	return iface.FileInfo{
		Size:         fi.Size(),
		LastModified: fi.ModTime(),
	}, nil
}
