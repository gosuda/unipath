package local

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gosuda.org/unipath/services"
)

type File struct {
	Object
	reader io.ReadCloser
}

func (f *File) ReadString() (string, error) {
	b, err := io.ReadAll(f)
	return string(b), err
}

func (f *File) WriteString(s string) error {
	b := []byte(s)
	n, err := f.Write(b)
	if n != len(b) {
		return errors.New(fmt.Sprintf("Wrote %v bytes from given %v bytes", n, len(b)))
	}
	return err
}

func (f *File) Directory() services.Directory {
	return &Directory{
		Object: Object{
			Path: filepath.Dir(f.Path),
		},
	}
}

func (f *File) Read(p []byte) (n int, err error) {
	if f.reader == nil {
		r, err := os.Open(f.Path)
		if err != nil {
			return 0, err
		}
		f.reader = r
	}

	return f.reader.Read(p)
}

func (f *File) Write(p []byte) (n int, err error) {
	file, err := os.Create(f.GetPath())
	n, err = file.Write(p)
	return n, err
}

func (f *File) Copy(newName string) error {
	return f.CopyTo(f.Directory(), newName)
}

func (f *File) CopyTo(dir services.Directory, newName ...string) error {
	var filename string
	if len(newName) > 0 {
		filename = newName[0]
	} else {
		_, filename = filepath.Split(f.GetPath())
	}

	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	file := dir.File(filename)
	n, err := file.Write(content)
	if err != nil {
		return err
	}
	if n != len(content) {
		return errors.New("couldn't copy full file")
	}

	return nil
}

func (f *File) Close() error {
	return f.reader.Close()
}
