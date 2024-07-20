package main

import (
	"flag"
	"fmt"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app"
	"os"
	"strconv"
)

const (
	grpcPort     = 50051
	httpPort     = 8081
	databaseUrl  = "postgres://test:test@localhost:5432/test"
	topic        = "loms.order-events"
	kafkaBrokers = "localhost:9092"
)

var opts = app.Options{}

func initOpts() {
	flag.StringVar(&opts.Addr, "addr", strconv.Itoa(httpPort), fmt.Sprintf("server address, default: %d", httpPort))
	flag.StringVar(&opts.GrpcAddr, "grpc-addr", strconv.Itoa(grpcPort), fmt.Sprintf("grpc address: %d", grpcPort))

	dbConnStr := os.Getenv("DATABASE_URL")
	if os.Getenv("DATABASE_URL") == "" {
		dbConnStr = databaseUrl
	}
	flag.StringVar(&opts.DbConnStr, "db-conn-str", dbConnStr, "database connection string")

	kafkaBrokersVar := os.Getenv("KAFKA_BROKERS")
	if os.Getenv("KAFKA_BROKERS") == "" {
		kafkaBrokersVar = kafkaBrokers
	}
	flag.StringVar(&opts.Brokers, "brokers", kafkaBrokersVar, "kafka brokers string")

	flag.StringVar(&opts.Topic, "topic", topic, "topic name")
	flag.Parse()
}
