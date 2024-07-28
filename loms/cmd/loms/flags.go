package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app"
)

const (
	grpcPort     = 50051
	httpPort     = 8081
	topic        = "loms.order-events"
	kafkaBrokers = "localhost:9092"
	dbURL1       = "postgresql://ws-8-user-1:ws-8-pass-1@localhost:8431/ws-8-db-1"
	dbURL2       = "postgresql://ws-8-user-2:ws-8-pass-2@localhost:8432/ws-8-db-2"
)

var opts = app.Options{}

func initOpts() {
	flag.StringVar(&opts.Addr, "addr", strconv.Itoa(httpPort), fmt.Sprintf("server address, default: %d", httpPort))
	flag.StringVar(&opts.GrpcAddr, "grpc-addr", strconv.Itoa(grpcPort), fmt.Sprintf("grpc address: %d", grpcPort))

	kafkaBrokersVar := os.Getenv("KAFKA_BROKERS")
	if kafkaBrokersVar == "" {
		kafkaBrokersVar = kafkaBrokers
	}
	flag.StringVar(&opts.Brokers, "brokers", kafkaBrokersVar, "kafka brokers string")

	flag.StringVar(&opts.Topic, "topic", topic, "topic name")

	dbUrl1 := os.Getenv("DOCKER_POSTGRES_DB_URL_1")
	if dbUrl1 == "" {
		dbUrl1 = dbURL1
	}

	flag.StringVar(&opts.DbUrl1, "db_url_1", dbUrl1, "postgres database url 1")

	dbUrl2 := os.Getenv("DOCKER_POSTGRES_DB_URL_2")
	if dbUrl2 == "" {
		dbUrl2 = dbURL2
	}
	flag.StringVar(&opts.DbUrl2, "db_url_2", dbUrl2, "postgres database url 2")

	flag.Parse()
}
