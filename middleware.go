package main

import (
	"net/http"
	"strings"
	"time"

	"go.mattglei.ch/timber"
)

// wrappedWriter provides a custom interface that allows us to store the status code of a request
// when it is being handled by our mux
type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.statusCode = code
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		timber.Donef(
			"%d [%s] %s %s",
			wrapped.statusCode,
			strings.ToLower(http.StatusText(wrapped.statusCode)),
			r.URL.Path,
			time.Since(start),
		)
	})
}
