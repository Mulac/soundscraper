package storage

import "os"

type Drive interface {
	SaveFile(file os.File) error
}

func NewDrive() (Drive, error) {
	// TODO(calum): turn into a factory.. for now just return GDrive impl
	return nil, nil
}