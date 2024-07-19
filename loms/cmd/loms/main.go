package main

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app"
	"log"
)

func main() {
	initOpts()

	ctx := context.Background()
	lomsApp, err := app.NewApp(ctx, app.NewConfig(opts))
	if err != nil {
		log.Fatalln(err)
	}

	lomsApp.ServeGrpcServer()

	err = lomsApp.CreateGatewayClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = lomsApp.GatewayListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
