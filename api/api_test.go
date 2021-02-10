package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alexdebril/feed-io-api-websocket/messaging"
)

func TestApi_ServeHTTP(t *testing.T) {
	buf := strings.NewReader("{\"title\": \"test\", \"feed_url\": \"http://localhost\"}")
	req := httptest.NewRequest(http.MethodPost, "/item", buf)
	writer := &responseWriter{
		expectedStatus: 204,
		testing:        t,
	}
	dispatcher := &testDispatcher{expectedCalls: 1}
	api := &Api{dispatcher}
	api.ServeHTTP(writer, req)

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

type testDispatcher struct {
	expectedCalls int
}

func (t *testDispatcher) Handle(item messaging.Item) {
	t.expectedCalls = t.expectedCalls + 1
}
