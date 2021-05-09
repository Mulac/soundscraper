package main

import (
	"fmt"
	"github.com/Mulac/soundscraper/storage"
	"os"
)

func main() {
	fmt.Println("hello")
	drive, err := storage.NewDrive()
	if err != nil {
		panic(err)
	}

	f, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	err = drive.SaveFile(f)
	if err != nil {
		panic(err)
	}
}
