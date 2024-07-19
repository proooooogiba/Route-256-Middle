package main

import (
	"flag"
	"fmt"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app"
	"strconv"
)

const (
	grpcPort  = 50051
	httpPort  = 8081
	dbConnStr = "postgres://test:test@localhost:5432/test?sslmode=disable"
	topic     = "loms.order-events"
)

var brokers = []string{"localhost:9092"}

var opts = app.Options{}

func initOpts() {
	flag.StringVar(&opts.Addr, "addr", strconv.Itoa(httpPort), fmt.Sprintf("server address, default: %d", httpPort))
	flag.StringVar(&opts.GrpcAddr, "grpc-addr", strconv.Itoa(grpcPort), fmt.Sprintf("grpc address: %d", grpcPort))
	flag.StringVar(&opts.DbConnStr, "db-conn-str", dbConnStr, "database connection string")
	flag.StringVar(&opts.Topic, "topic", topic, "topic name")
	flag.Parse()
}
