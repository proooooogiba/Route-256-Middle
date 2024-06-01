package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("New Request method: %s remote_addr: %s url: %s time: %s", r.Method, r.RemoteAddr, r.URL, time.Since(start).String())
	})
}
