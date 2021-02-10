package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexdebril/feed-io-api-websocket/messaging"
)

type Api struct {
	messaging.Dispatcher
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}
	item, err := parseItem(r)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(500)
	}
	a.Dispatcher.Handle(item)
	w.WriteHeader(204)
}

func parseItem(r *http.Request) (messaging.Item, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return messaging.Item{}, err
	}
	i := messaging.Item{}
	err = json.Unmarshal(body, &i)
	return i, err
}
