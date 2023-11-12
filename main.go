package main

import (
	"log"
	"net/http"

	"http-request-canceled/handler"
)

var port = ":8080"

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", handler.HandleIndex)
	mux.HandleFunc("/post", handler.HandlePost)
	mux.HandleFunc("/custom", handler.HandleCustom)

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)

	log.Printf("server running on localhost%s\n", port)

	server := new(http.Server)
	server.Addr = port
	server.Handler = handler
	server.ListenAndServe()
}
