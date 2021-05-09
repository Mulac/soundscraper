package main

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func Download(videoID string) {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	resp, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}
