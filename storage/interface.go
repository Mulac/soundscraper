package storage

import (
	"os"
)

type Drive interface {
	SaveFile(file *os.File) error
}

//func NewDrive() (Drive, error) {
//	// TODO(calum): turn into a factory.. for now just return GDrive impl
//	drive, err := newGoogleDrive()
//	if err != nil {
//		return nil, fmt.Errorf("ERROR|NewDrive()|%v", err)
//	}
//	return drive, nil
//}