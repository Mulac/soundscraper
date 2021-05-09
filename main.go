package main

import (
	"fmt"
	"github.com/Mulac/soundscraper/storage"
	"github.com/Mulac/soundscraper/scraper"
)

func main() {
	fmt.Println("hello")
	_ = storage.GoogleDrive{}

	videoID := "NsKaCS3CtsY"
	scraper.Download(videoID)
}
