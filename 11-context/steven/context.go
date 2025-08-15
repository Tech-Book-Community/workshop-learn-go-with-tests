package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// 建立 HTTP handler，從 store 取得資料，然後直接寫到 HTTP 回應
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context()) //帶入 context 的參數

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
