package storage

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Mulac/soundscraper/data"

	"golang.org/x/oauth2/google"
	gDrive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const FOLDER = "1cibm15gWSGs6bNWSnop5CGTp2HlIJ7EI"

type googleDrive struct {
	service *gDrive.Service
}

func (gd *googleDrive) SaveFile(file data.File) error {
	f := &gDrive.File{
		Name:    file.Name(),
		Parents: []string{FOLDER},
	}
	_, err := gd.service.Files.Create(f).Media(file).Do()
	if err != nil {
		return fmt.Errorf("ERROR|Drive/googleDrive|SaveFile(%s)|could not save file|%v", file.Name(), err)
	}

	return nil
}

func newGoogleDrive() (*googleDrive, error) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, fmt.Errorf("ERROR|storage/googleDrive|newGoogleDrive()|unable to read client secret file|%v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gDrive.DriveScope, gDrive.DriveFileScope, gDrive.DriveAppdataScope)
	if err != nil {
		return nil, fmt.Errorf("ERROR|storage/googleDrive|newGoogleDrive()|unable to parse client secret file to config|%v", err)
	}

	srv, err := gDrive.NewService(context.Background(), option.WithHTTPClient(getClient(config)))
	if err != nil {
		return nil, fmt.Errorf("ERROR|storage/googleDrive|newGoogleDrive()|unable to retrieve Drive client|%v", err)
	}

	return &googleDrive{
		service: srv,
	}, nil
}
