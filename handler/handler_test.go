package handler

import (
	"net/http"
	"testing"
)

func TestItemHandler_ServeHTTP(t *testing.T) {
	req := &http.Request{
		Method: "POST",
	}
	writer := &responseWriter{
		expectedStatus: 204,
		testing:        t,
	}
	h := &ItemHandler{}
	h.ServeHTTP(writer, req)
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
