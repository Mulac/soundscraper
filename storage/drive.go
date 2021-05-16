package storage

import (
	"fmt"
	"github.com/Mulac/soundscraper/config"
	"github.com/Mulac/soundscraper/util"
	"log"
	"sync"
)

type DriveType string

const (
	DRIVE_NAIVE  DriveType = "naive"
	DRIVE_GOODLE DriveType = "googleDrive"
)

var driveSingleton DriveManager
var once sync.Once

// Drive returns the default DriveManager implementation set via config
func Drive() DriveManager {
	once.Do(func() {
		d, err := NewDriveFactory().SetType(DriveType(config.Manager().GetString(util.ENV_DRIVE_IMPL))).New()
		if err != nil {
			log.Fatal(err)
		}
		driveSingleton = d
	})

	return driveSingleton
}

type driveFactory struct {
	dType DriveType
}

func (df *driveFactory) SetType(dType DriveType) DriveFactory {
	df.dType = dType
	return df
}

func (df *driveFactory) New() (DriveManager, error) {
	switch df.dType {
	case DRIVE_NAIVE, "":
		return newNaiveDriveManager()
	case DRIVE_GOODLE:
		return newGoogleDrive()
	default:
		return nil, fmt.Errorf("ERROR|DriveFactory|drive type %s not recognised", df.dType)
	}
}

func NewDriveFactory() DriveFactory {
	return &driveFactory{}
}
