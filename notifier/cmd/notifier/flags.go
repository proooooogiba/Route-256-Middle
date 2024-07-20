package main

import (
	"flag"
	"os"
)

type flags struct {
	topic             string
	bootstrapServer   string
	consumerGroupName string
}

func init() {
	flag.StringVar(&cliFlags.topic, "topic", "loms.order-events", "topic to produce")
	bootstrapServer := os.Getenv("KAFKA_HOST")
	if bootstrapServer == "" {
		bootstrapServer = "localhost:9092"
	}

	flag.StringVar(&cliFlags.bootstrapServer, "bootstrap-server", bootstrapServer, "kafka broker host and port")
	flag.StringVar(&cliFlags.consumerGroupName, "notifier-consumer-group", "notifier-consumer-group", "the consumer group name, which is reading messages")

	flag.Parse()
}
