package main

import (
	"context"
	"log"

	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app"
)

func main() {
	initOpts()

	log.Println("loms app strating..")

	ctx := context.Background()
	lomsApp, err := app.NewApp(ctx, app.NewConfig(opts))
	if err != nil {
		log.Fatalln(err)
	}
	defer lomsApp.Close()

	log.Println("loms app started")

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
