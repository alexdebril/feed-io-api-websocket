package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexdebril/feed-io-api-websocket/handler"
	"io/ioutil"
	"log"
	"net/http"
)

var ItemChannel chan handler.Item

func main() {
	log.Println("starting server ...")
	mux := http.NewServeMux()

	mux.HandleFunc("/item", ReceiveMessage)
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

func ReceiveMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}
	item, err := parseItem(r)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(500)
	}
	ItemChannel <- item
	w.WriteHeader(204)
}

func parseItem(r *http.Request) (handler.Item, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return handler.Item{}, err
	}
	i := handler.Item{}
	err = json.Unmarshal(body, &i)
	return i, err
}

func toJson(item handler.Item) string {
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

	ItemChannel = make(chan handler.Item)
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
