package main

import (
	"fmt"
	"os"

	"github.com/Mulac/soundscraper/scraper"
	"github.com/Mulac/soundscraper/storage"
)

func main() {
	fmt.Println("hello")

	id := "YQHsXMglC9A"
	scraper.Download().Download(id)

	f, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	err = storage.Drive().SaveFile(f)
	if err != nil {
		panic(err)
	}
}
