package context

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	t         *testing.T
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	dataChan := make(chan string)

	go func() {
		var result string
		for _, v := range s.response {
			select {
			case <-ctx.Done():
				fmt.Println("spy got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(v)
			}
		}

		dataChan <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-dataChan:
		return result, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	fmt.Println("natawag header")
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write(b []byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, word"
		store := &SpyStore{
			response: data,
			t:        t,
		}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		if got := res.Body.String(); got != data {
			t.Errorf("got %q, want %q", got, data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{
			response: data,
			t:        t,
		}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(ctx)
		spyRes := &SpyResponseWriter{}

		server.ServeHTTP(spyRes, req)

		if spyRes.written {
			t.Error("a response should not have been written")
		}
	})
}
