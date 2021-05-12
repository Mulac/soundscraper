package scraper

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/kkdai/youtube/v2"
)

type Scraper interface {
	Download()
}

type scraperImpl struct {
}

func (s scraperImpl) Download(videoID string) {
	for i := 0; i < 22; i++ {
		client := youtube.Client{}

		video, err := client.GetVideo(videoID)
		if err != nil {
			print("error")
		}

		resp, err := client.GetStream(video, &video.Formats[i])

		if err != nil {
			print("error")
		}
		defer resp.Body.Close()

		videoname := "Format" + strconv.Itoa(i) + ".mp4"
		fmt.Printf("%s : %+v \n", videoname, video.Formats[i])
		file, err := os.Create(videoname)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
	}
}
