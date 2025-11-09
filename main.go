package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Initializing server...")
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	server.ListenAndServe()
}

func healthHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "text/plain; charset=utf-8")
	responseWriter.WriteHeader(200)
	responseWriter.Write([]byte("OK"))
}
