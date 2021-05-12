package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Mulac/soundscraper/config"
	"github.com/Mulac/soundscraper/util"
)

// naiveDriveManager stores files in a local directory
// by default the directory where they are stored is 'soundscraperdrive' in your home directory
type naiveDriveManager struct {
	root string
}

func (dm *naiveDriveManager) SaveFile(file *os.File) error {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ERROR|Drive/naive|SaveFile(%s)|unable to read contents from file|%v", file.Name(), err)
	}

	err = os.WriteFile(filepath.Join(dm.root, file.Name()), content, 0666)
	if err != nil {
		return fmt.Errorf("ERROR|Drive/naive|SaveFile(%s)|unable to write file to %s|%v", file.Name(), dm.root, err)
	}
	return nil
}

func newNaiveDriveManager() (*naiveDriveManager, error) {
	root := config.Manager().GetString(util.ENV_DRIVE_ROOT)
	if root == "" {
		// if no directory is set in config create a new directory called 'soundscraperdrive' in home directory
		home, _ := os.UserHomeDir()
		root = filepath.Join(home, "soundscraperdrive")
		if err := os.Mkdir(root, 0777); err != nil && os.IsNotExist(err) {
			return nil, fmt.Errorf("can't creat root dir %s|%T", root, err)
		}
	}
	return &naiveDriveManager{root}, nil
}
