package main

import (
	"fmt"
	"net/http"

	"github.com/Mulac/soundscraper/scraper"
)

func downloadYoutubeVideo(w http.ResponseWriter, r *http.Request) {
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
