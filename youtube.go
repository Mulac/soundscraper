package main

import (
	"fmt"
	"net/http"
)

func downloadYoutubeVideo(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == "" {
		// TODO
	}

	// TODO(calum): download video
	fmt.Fprintf(w, "Hello!")
}
