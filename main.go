package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Initializing server...")
	port := ":8080"
	mux := http.NewServeMux()
	fileServerHandler := http.FileServer(http.Dir("."))
	server := &http.Server{
		Addr: port,
		Handler: mux,
	}
	mux.Handle("/", fileServerHandler)
	server.ListenAndServe()
}
