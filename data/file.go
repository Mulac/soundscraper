package data

import (
	"fmt"
	"io"
)

type File interface {
	io.Reader
	Name() string
}

type fileImpl struct {
	io.Reader
	FileName string
}

func (f *fileImpl) Name() string {
	return f.FileName
}

func NewFile(name string, content *io.Reader) (File, error) {
	if name == "" {
		return nil, fmt.Errorf("NewFile()|cannot create file with no name")
	}
	return &fileImpl{
		Reader: *content,
		FileName: name,
	}, nil
}
