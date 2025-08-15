package myserver

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response   string
	fetchSleep time.Duration
	t          *testing.T
}

var _ Store = (*SpyStore)(nil)

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			// 假裝每個字都很慢
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(s.fetchSleep)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case d := <-data:
		return d, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, fetchSleep: time.Millisecond, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {

		data := "hello world"

		cancelDuration := 1 * time.Millisecond
		fetchSleep := 2 * time.Millisecond
		store := &SpyStore{
			response:   data,
			fetchSleep: fetchSleep,
			t:          t,
		}

		server := Server(store)

		ctx, cancel := context.WithCancel(context.Background())

		time.AfterFunc(cancelDuration, cancel)
		req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		time.Sleep(cancelDuration)

		if res.Body.String() != "" {
			t.Errorf(`got "%s", want "%s"`, res.Body.String(), "")
		}
	})
}
