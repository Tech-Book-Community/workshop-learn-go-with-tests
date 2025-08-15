package myserver

import (
	"errors"
	"net/http"
)

type SpyResponseWriter struct {
	written bool
}

var _ http.ResponseWriter = (*SpyResponseWriter)(nil)

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
