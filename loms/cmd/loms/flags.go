package main

import (
	"strings"
)

const (
	grpcPort  = 50051
	httpPort  = 8081
	dbConnStr = "postgres://test:test@localhost:5432/test?sslmode=disable"
)

func headerMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case "x-auth":
		return key, true
	default:
		return key, false
	}
}
