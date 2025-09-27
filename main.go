package main

import (
	"log"
	"net/http"

	"tailscale.com/tsnet"
)

func main() {
	srv := &tsnet.Server{
		Hostname: getEnv("TS_HOSTNAME", "paperless-share"),
		Dir:      getEnv("TS_DIR", "./tsnet-data"),
		AuthKey:  getEnv("TS_AUTHKEY", ""),
	}
	defer srv.Close()

	ln, err := srv.ListenFunnel("tcp", ":443")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRoot(w, r)
	})
	http.HandleFunc("/share/", handleShare)

	// start the server
	err = http.Serve(ln, nil)
	if err != nil {
		log.Fatal(err)
	}
}
