package scraper

import (
	"fmt"
	"strconv"

	"github.com/Mulac/soundscraper/data"
	"github.com/Mulac/soundscraper/storage"
	"github.com/kkdai/youtube/v2"
)

type youtubeScraper struct {
	client youtube.Client
}

func newYoutubeDownloadManager() (*youtubeScraper, error) {
	client := youtube.Client{}
	return &youtubeScraper{client: client}, nil

}

func (s *youtubeScraper) Download(videoID string) error {
	video, err := s.client.GetVideo(videoID)
	if err != nil {
		print("error")
	}

	resp, err := s.client.GetStream(video, &video.Formats[0])

	if err != nil {
		print("error")
	}
	defer resp.Body.Close()

	videoname := "Format" + strconv.Itoa(0) + ".mp4"
	fmt.Printf("%s : %+v \n", videoname, video.Formats[0])

	// TODO(lucy): instead of storing the file to disk - use the storagemanager
	file, err := data.NewFile(videoname, resp.Body)
	if err != nil {
		return fmt.Errorf("ERROR|scraper/youtube|Download()|unable to find file|%v", err)
	}
	storage.Drive().SaveFile(file)
	if err != nil {
		return fmt.Errorf("ERROR|scraper/youtube|Download()|error in saving the file|%v", err)
	}

	return nil
}
