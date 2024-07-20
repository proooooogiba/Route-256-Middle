package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app"
)

const (
	grpcPort    = 50051
	httpPort    = 8081
	databaseUrl = "postgres://test:test@postgres-loms:5432/test"
	topic       = "loms.order-events"
)

var opts = app.Options{}

func initOpts() {
	flag.StringVar(&opts.Addr, "addr", strconv.Itoa(httpPort), fmt.Sprintf("server address, default: %d", httpPort))
	flag.StringVar(&opts.GrpcAddr, "grpc-addr", strconv.Itoa(grpcPort), fmt.Sprintf("grpc address: %d", grpcPort))
	flag.StringVar(&opts.DbConnStr, "db-conn-str", os.Getenv("DATABASE_URL"), "database connection string")
	flag.StringVar(&opts.Topic, "topic", topic, "topic name")
	flag.Parse()
}
