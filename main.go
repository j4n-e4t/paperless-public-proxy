package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	baseURL := getEnv("PAPERLESS_BASE_URL", "http://paperless:8000/")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	
	http.HandleFunc("/share/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/share/")

		resp, err := http.Get(baseURL + "share/" + id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK || resp.Header.Get("Content-Type") != "application/pdf" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		io.Copy(w, resp.Body)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
