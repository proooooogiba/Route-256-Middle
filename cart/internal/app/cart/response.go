package cart

import (
	"log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	if err != nil {
		log.Printf("Error writing response: %s\n", err)
	}
}
