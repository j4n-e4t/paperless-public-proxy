package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"tailscale.com/tsnet"
)

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

	// Copy headers
	w.Header().Set("Content-Type", "application/pdf")
	io.Copy(w, resp.Body)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	srv := &tsnet.Server{
		Hostname: "paperless-share",
		Dir: "./ts",
	}
	ln, err := srv.ListenFunnel("tcp", ":443")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Listening on the Tailscale Funnel Hostname")

	http.HandleFunc("/", notFound)
	http.HandleFunc("/share/", handleShare)

	// start the server	
	err = http.Serve(ln, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
