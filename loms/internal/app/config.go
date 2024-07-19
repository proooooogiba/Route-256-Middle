package app

type (
	Options struct {
		Addr, DbConnStr, GrpcAddr, Topic string
	}

	path struct {
		string
	}
	config struct {
		grpcPort  string
		httpPort  string
		topic     string
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
		path:      path{},
	}
}
