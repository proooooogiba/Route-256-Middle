package mw

import (
	"log"
	"net/http"
)

func WithHTTPLoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI, r.Proto, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
