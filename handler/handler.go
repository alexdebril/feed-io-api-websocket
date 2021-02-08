package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Title string `json:"title"`
	FeedUrl string `json:"feed_url"`
}

type ItemHandler struct {
	Message chan Item
}

func (h *ItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}
	item, err := h.getItem(r)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(500)
	}
	log.Println("yo")
	h.Message <- item
	log.Println("yo man")
	w.WriteHeader(204)
}

func (h *ItemHandler) getItem(r *http.Request) (Item, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Item{}, err
	}
	i := Item{}
	err = json.Unmarshal(body, &i)
	return i, err
}
