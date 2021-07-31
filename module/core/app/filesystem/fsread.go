package filesystem

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func (f *FileSystem) PeekFile(path string, maxSize int) ([]byte, error) {
	var handleError = func(err error) ([]byte, error) {
		for _, c := range f.children {
			if x, err := c.PeekFile(path, maxSize); err == nil {
				return x, nil
			}
		}
		return nil, errors.Wrapf(err, "unable to read file [%s]", path)
	}

	file, err := os.Open(f.getPath(path))
	if err != nil {
		return handleError(err)
	}
	defer func() { _ = file.Close() }()

	b := make([]byte, maxSize)

	size, err := file.Read(b)
	if err != nil {
		return handleError(err)
	}
	return b[:size], nil
}

func (f *FileSystem) ReadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(f.getPath(path))
	if err != nil {
		for _, c := range f.children {
			if x, err := c.ReadFile(path); err == nil {
				return x, nil
			}
		}
		return nil, errors.Wrapf(err, "unable to read file [%s]", path)
	}
	return b, nil
}
