package main

import (
	"github.com/alexdebril/feed-io-api-websocket/websocket"
	"log"
	"net/http"

	"github.com/alexdebril/feed-io-api-websocket/api"
	"github.com/alexdebril/feed-io-api-websocket/messaging"
)

func main() {
	log.Println("starting server ...")
	mux := http.NewServeMux()
	dispatcher := messaging.NewDispatcher(256)

	api := &api.Api{
		Dispatcher: dispatcher,
	}
	ws := &websocket.Websocket{
		Dispatcher: dispatcher,
	}
	mux.Handle("/item", api)
	mux.Handle("/read", ws)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("impossible to start the HTTP server")
	}

}
