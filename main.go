package main

// Static content example page with Hello world

import (
	"io"
	"log"
	"net/http"
)

const version string = "2.0.3"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number really
func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

func main() {
	log.Printf("Listening on port 8000....")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/version", versionHandler)
	http.ListenAndServe(":8000", nil)
}

