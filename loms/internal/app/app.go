package app

import "net/http"

type App struct {
	mux    *http.ServeMux
	config config
	server *http.Server
}

func NewApp(config config) (*App, error) {
	mux := http.NewServeMux()
	return &App{
		mux:    mux,
		config: config,
		server: &http.Server{Addr: config.addr, Handler: InitMiddlewares(mux)},
	}, nil
}

func (a *App) ListenAndServe() error {
	return a.server.ListenAndServe()
}

func (a *App) Close() error {
	return a.server.Close()
}

func InitMiddlewares(router *http.ServeMux) http.Handler {
	return router
}
