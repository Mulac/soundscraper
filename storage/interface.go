package storage

import (
	"os"
)

// DriveManager is a facade for interacting with some kind of file store.
//
// The purpose of a facade is to hide the underlying implementation from the user,
// this is most likely other developers, so that they can get on with their job.
//
// The naive.go implementation simply stores the files in a directory on your local machine,
// we may need to store files elsewhere though, for example GoogleDrive, a network file system
// or some other database like mongoDB.
type DriveManager interface {
	SaveFile(file *os.File) error
}

// DriveFactory is how we can create new DriveManagers
// encapsulating all the logic needed to spin up specific implementations
type DriveFactory interface {
	SetType(dType DriveType) DriveFactory
	New() (DriveManager, error)
}
