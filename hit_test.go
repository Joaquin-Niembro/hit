package hit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestSendFunction(t *testing.T) {
	t.Parallel()
	_, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatalf("creating http request: %v", err)
	}

}

func TestSendNFunction(t *testing.T) {
	t.Parallel()
	var hits atomic.Int64
	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		hits.Add(1)
	}))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, http.NoBody)
	if err != nil {
		t.Fatalf("creating http request: %v", err)
	}
	results, err := SendN(t.Context(), 10, req, Options{Concurrency: 5})
	if err != nil {
		t.Fatalf("SendN() err=%v, want nil", err)
	}
	for range results {
		fmt.Println("Consuming results")
	}
	if got := hits.Load(); got != 10 {
		t.Errorf("hits.Load() = %v, want 10", got)
	}
}
