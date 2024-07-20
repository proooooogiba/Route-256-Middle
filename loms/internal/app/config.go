package app

import "strings"

type (
	Options struct {
		Addr, DbConnStr, GrpcAddr, Topic, Brokers string
	}

	path struct {
		string
	}
	config struct {
		grpcPort  string
		httpPort  string
		topic     string
		brokers   []string
		dbConnStr string
		path      path
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
	}
}
