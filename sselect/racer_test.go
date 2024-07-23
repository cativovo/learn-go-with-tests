package sselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Millisecond * 100)
		defer slowServer.Close()
		fastServer := makeDelayedServer(time.Millisecond * 10)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("didn't expect an error but got one %v", err)
		}

		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(time.Millisecond * 20)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, time.Millisecond*10)
		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
