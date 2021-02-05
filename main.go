package main

import (
	"github.com/alexdebril/feed-io-api-websocket/handler"
	"log"
	"net/http"
)

func main() {
	log.Println("starting server ...")
	mux := http.NewServeMux()
	h := &handler.ItemHandler{}

	mux.Handle("/item", h)
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("impossible to start the HTTP server")
	}
}
