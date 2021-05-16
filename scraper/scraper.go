package scraper

import (
	"fmt"
	"log"
	"sync"

	"github.com/Mulac/soundscraper/config"
	"github.com/Mulac/soundscraper/util"
)

type DownloadType string

const (
	DOWNLOAD_YOUTUBE DownloadType = "youtube"
)

var downloadSingleton DownloadManager
var once sync.Once

func Download() DownloadManager {
	once.Do(func() {
		d, err := NewDownloadFactory().SetType(DownloadType(config.Manager().GetString(util.ENV_DOWNLOAD_IMPL))).New()
		if err != nil {
			log.Fatal(err)
		}
		downloadSingleton = d
	})

	return downloadSingleton
}

type downloadFactory struct {
	dtype DownloadType
}

func (df *downloadFactory) SetType(dType DownloadType) DownloadFactory {
	df.dtype = dType
	return df
}

func (df *downloadFactory) New() (DownloadManager, error) {
	switch df.dtype {
	case DOWNLOAD_YOUTUBE, "":
		return newYoutubeDownloadManager()

	default:
		return nil, fmt.Errorf("ERROR|DownloadFactory|download type %s not recognised", df.dtype)
	}
}

func NewDownloadFactory() DownloadFactory {
	return &downloadFactory{}
}
