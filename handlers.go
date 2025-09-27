package main

import (
	"io"
	"net/http"
	"strings"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func handleShare(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/share/")

	resp, err := http.Get(getEnv("PAPERLESS_BASE_URL", "http://paperless:8000/") + "share/" + id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK ||
		resp.Header.Get("Content-Type") != "application/pdf" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	io.Copy(w, resp.Body)
}
