package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestItemHandler_ServeHTTP(t *testing.T) {
	buf := strings.NewReader("{\"title\": \"test\", \"feed_url\": \"http://localhost\"}")
	req := httptest.NewRequest(http.MethodPost, "/item", buf)
	writer := &responseWriter{
		expectedStatus: 204,
		testing:        t,
	}
	var msg chan Item
	msg = make(chan Item)
	defer func() {
		close(msg)
		msg = nil
	}()
	h := &ItemHandler{
		Message: msg,
	}
	h.ServeHTTP(writer, req)
	message := <-msg
	log.Printf("%v", message)
}

type responseWriter struct {
	expectedStatus int
	testing        *testing.T
}

func (r *responseWriter) Header() http.Header {
	return nil
}

func (r *responseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (r *responseWriter) WriteHeader(statusCode int) {
	if statusCode != r.expectedStatus {
		r.testing.Fatalf("unexpected status: %v", statusCode)
	}
}
