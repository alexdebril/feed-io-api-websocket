package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Title string `json:"title"`
	FeedUrl string `json:"feed_url"`
}

type ItemHandler struct {}

func (h *ItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(204)
}

func (h *ItemHandler) getItem(r *http.Request) (*Item, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	i := &Item{}
	err = json.Unmarshal(body, i)
	return i, err
}
