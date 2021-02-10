package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexdebril/feed-io-api-websocket/api"
	"github.com/alexdebril/feed-io-api-websocket/messaging"
)

var ItemChannel chan messaging.Item

func main() {
	log.Println("starting server ...")
	mux := http.NewServeMux()
	dispatcher := &ChannelDispatcher{}

	api := &api.Api{
		Dispatcher: dispatcher,
	}
	mux.Handle("/item", api)
	mux.HandleFunc("/read", publish)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("impossible to start the HTTP server")
	}
}

type ChannelDispatcher struct{}

func (d *ChannelDispatcher) Handle(item messaging.Item) {
	ItemChannel <- item
}

func toJson(item messaging.Item) string {
	out, err := json.Marshal(item)
	if err != nil {
		return "{}"
	}
	return string(out)
}

func publish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ItemChannel = make(chan messaging.Item)
	defer func() {
		close(ItemChannel)
		ItemChannel = nil
		log.Printf("client connection is closed")
	}()

	flusher, _ := w.(http.Flusher)

	for {
		select {
		case item := <-ItemChannel:
			_, _ = fmt.Fprintf(w, "%s\n", toJson(item))
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
