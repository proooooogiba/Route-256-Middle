package main

import (
	"github.com/pkg/errors"
	"log"
	"net/http"
	"route256/cart/internal/app"
)

func main() {
	initOpts()
	service, err := app.NewApp(app.NewConfig(opts))
	if err != nil {
		log.Fatal("{FATAL} ", err)
	}

	err = service.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("server closed\n")
	}
	if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
	return
}
