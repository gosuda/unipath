package local

import (
	"os"
	"strings"

	"gosuda.org/unipath/transfer/iface"
)

type Directory struct {
	Object
}

func (d *Directory) File(path string) iface.File {
	return &File{
		Object: Object{
			Path: d.Path + "/" + path,
		},
		reader: nil,
	}
}

func (d *Directory) Directory(path string) iface.Directory {
	path = d.Path + "/" + path
	path = strings.TrimRight(path, "/")
	return &Directory{
		Object{
			Path: path,
		},
	}
}

func (d *Directory) Create() error {
	return os.Mkdir(d.Path, os.ModePerm)
}

func (d *Directory) Files() ([]iface.File, error) {
	files, err := os.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}
	var localFiles []iface.File
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		localFile := File{
			Object: Object{
				Path: d.Path + "/" + file.Name(),
			},
			reader: nil,
		}
		localFiles = append(localFiles, iface.File(&localFile))
	}
	return localFiles, nil
}

func (d *Directory) Directories() ([]iface.Directory, error) {
	files, err := os.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}
	var localDirs []iface.Directory
	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}

		localDir := Directory{
			Object: Object{
				Path: d.Path + "/" + dir.Name(),
			},
		}
		localDirs = append(localDirs, iface.Directory(&localDir))
	}
	return localDirs, nil
}
