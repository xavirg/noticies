package generator

import (
	"io"
	"os"
)

type FileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
	Create(name string) (io.WriteCloser, error)
}

type OSFileSystem struct{}

func (OSFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (OSFileSystem) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}
