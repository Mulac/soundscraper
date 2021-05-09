package storage

import "os"

type Drive interface {
	SaveFile(file os.File) error
}