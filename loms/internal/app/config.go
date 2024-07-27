package app

import "strings"

type (
	Options struct {
		Addr, DbConnStr, GrpcAddr, Topic, Brokers string
	}

	database struct {
		DSN string
	}

	path struct {
		string
	}
	config struct {
		grpcPort     string
		httpPort     string
		topic        string
		brokers      []string
		dbConnStr    string
		path         path
		databasePool []database
	}
)

func NewConfig(opts Options) config {
	return config{
		dbConnStr: opts.DbConnStr,
		grpcPort:  opts.GrpcAddr,
		httpPort:  opts.Addr,
		topic:     opts.Topic,
		brokers:   strings.Split(opts.Brokers, ","),
		path:      path{},
		databasePool: []database{
			// индекс это номер шарда,
			// индексы важны и их нельзя менять
			0: {
				DSN: "postgresql://ws-8-user-1:ws-8-pass-1@localhost:8431/ws-8-db-1",
			},
			1: {
				DSN: "postgresql://ws-8-user-2:ws-8-pass-2@localhost:8432/ws-8-db-2",
			},
		},
	}
}
