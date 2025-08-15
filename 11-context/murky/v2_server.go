package myserver

import (
	"context"
	"fmt"
	"net/http"
)

// Store is store
type Store interface {
	Fetch(context.Context) (string, error)
}

// Server create http.HandlerFunc
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}
