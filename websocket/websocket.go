package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alexdebril/feed-io-api-websocket/messaging"
)

type Websocket struct {
	messaging.Dispatcher
}

func (ws *Websocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id, channel := ws.Dispatcher.GetChannel()
	defer ws.Dispatcher.Release(id)
	flusher, _ := w.(http.Flusher)
	start := time.Now().String()
	_, _ = fmt.Fprintf(w, "event: handshake\ndata: {\"date\": \"%s\"}\n\n", start)
	flusher.Flush()
	for {
		select {
		case item := <-channel:
			_, _ = fmt.Fprintf(w, "event: item\ndata: %s\n\n", toJson(item))
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

func toJson(item messaging.Item) string {
	out, err := json.Marshal(item)
	if err != nil {
		return "{}"
	}
	return string(out)
}
