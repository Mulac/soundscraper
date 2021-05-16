package scraper

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/kkdai/youtube/v2"
)

type youtubeScraper struct {
}

func newYoutubeDownloadManager() (*youtubeScraper, error) {
	return &youtubeScraper{}, nil

}

func (s *youtubeScraper) Download(videoID string) error {
	// TODO(lucy): put the client in the youtubeScraper struct so that we don't create new client each time we download a video
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		print("error")
	}

	resp, err := client.GetStream(video, &video.Formats[0])

	if err != nil {
		print("error")
	}
	defer resp.Body.Close()

	videoname := "Format" + strconv.Itoa(0) + ".mp4"
	fmt.Printf("%s : %+v \n", videoname, video.Formats[0])

	// TODO(lucy): instead of storing the file to disk - use the storagemanager
	file, err := os.Create(videoname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	return nil
}
