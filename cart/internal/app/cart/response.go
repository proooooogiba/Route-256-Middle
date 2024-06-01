package cart

import (
	"go.uber.org/zap"
	"net/http"
)

func WriteResponse(logger *zap.SugaredLogger, w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	if err != nil {
		logger.Errorw("error in writing response body", "error", err)
	}
}
