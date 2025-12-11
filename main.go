package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRoot(w, r)
	})
	http.HandleFunc("/share/", handleShare)

	port := getEnv("PORT", ":7070")
	log.Printf("Starting server on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
