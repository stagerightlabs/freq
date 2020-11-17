package main

import (
	"fmt"
	"log"
	"net/http"
)

// startServer launches a web server to process API requests
func startServer(host, port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/freq", apiFreqHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", host, port),
		Handler: mux,
	}

	log.Printf("Listening on https://%v:%v", host, port)
	log.Fatal(server.ListenAndServe())
}
