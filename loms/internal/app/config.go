package app

type (
	Options struct {
		Addr string
	}

	path struct {
		string
	}
	config struct {
		addr string
		path path
	}
)

func NewConfig(opts Options) config {
	return config{
		addr: opts.Addr,
		path: path{},
	}
}
