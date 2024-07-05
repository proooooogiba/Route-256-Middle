package main

import (
	"strings"
)

const (
	grpcPort = 50051
	httpPort = 8081
)

func headerMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case "x-auth":
		return key, true
	default:
		return key, false
	}
}
