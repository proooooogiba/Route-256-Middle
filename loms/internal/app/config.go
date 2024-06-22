package app

type (
	Options struct {
		Addr, DbConnStr string
	}

	path struct {
		string
	}
	config struct {
		addr      string
		dbConnStr string
		path      path
	}
)

func NewConfig(opts Options) config {
	return config{
		addr:      opts.Addr,
		dbConnStr: opts.DbConnStr,
		path:      path{},
	}
}
