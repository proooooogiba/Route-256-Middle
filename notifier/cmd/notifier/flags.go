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
	flag.StringVar(&cliFlags.bootstrapServer, "bootstrap-server", os.Getenv("KAFKA_HOST"), "kafka broker host and port")
	flag.StringVar(&cliFlags.consumerGroupName, "notifier-consumer-group", "notifier-consumer-group", "the consumer group name, which is reading messages")

	flag.Parse()
}
