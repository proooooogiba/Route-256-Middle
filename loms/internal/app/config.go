package app

import (
	"strings"
)

type (
	Options struct {
		Addr, GrpcAddr, Topic, Brokers string
		DbUrl1, DbUrl2                 string
	}

	database struct {
		DSN string
	}

	path struct {
		string
	}
	Config struct {
		grpcPort     string
		httpPort     string
		topic        string
		brokers      []string
		path         path
		databasePool []database
	}
)

func NewConfig(opts Options) Config {
	return Config{
		grpcPort: opts.GrpcAddr,
		httpPort: opts.Addr,
		topic:    opts.Topic,
		brokers:  strings.Split(opts.Brokers, ","),
		path:     path{},
		databasePool: []database{
			0: {
				DSN: opts.DbUrl1,
			},
			1: {
				DSN: opts.DbUrl2,
			},
		},
	}
}
