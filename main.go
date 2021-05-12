package main

import (
	"fmt"
	"os"

	"github.com/Mulac/soundscraper/storage"
)

func main() {
	fmt.Println("hello")

	f, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	err = storage.Drive().SaveFile(f)
	if err != nil {
		panic(err)
	}
}
