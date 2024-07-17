package main

import (
	"strings"
)

const (
	grpcPort  = 50051
	httpPort  = 8081
	dbConnStr = "postgres://test:test@localhost:5432/test?sslmode=disable"
	topic     = "loms.order-events"
)

var brokers = []string{"localhost:9092"}

func headerMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case "x-auth":
		return key, true
	default:
		return key, false
	}
}
