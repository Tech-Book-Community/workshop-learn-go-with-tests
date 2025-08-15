package myserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServerV3(t *testing.T) {
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
    // 測試目的：當請求被取消時，伺服器不應該寫入任何回應
    t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
        data := "hello, world"
        store := &SpyStore{response: data, fetchSleep: 10 * time.Millisecond, t: t}
        svr := Server(store)

        request := httptest.NewRequest(http.MethodGet, "/", nil)

        cancellingCtx, cancel := context.WithCancel(request.Context())
        time.AfterFunc(5 * time.Millisecond, cancel)
        request = request.WithContext(cancellingCtx)

        response := &SpyResponseWriter{}

        svr.ServeHTTP(response, request)

        // 呼叫伺服器後，檢查 response.written 是否為 false
        // 也就是伺服器在請求被取消時，不應該寫入任何回應。
        if response.written {
            t.Error("a response should not have been written")
        }
    })
}
