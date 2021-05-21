package main

import (
	"fmt"
	"github.com/Mulac/soundscraper/scraper"
	"log"
	"net/http"
)

func DownloadYoutubeVideo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, fmt.Sprintf("could not download video: id empty"), 400)
		return
	}

	err := scraper.Download().Download(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not download video: %v", err), 400)
		return
	}

	// video downloaded successfully
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/youtube/video", DownloadYoutubeVideo)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
