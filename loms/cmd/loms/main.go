package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app/loms"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/mw"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/producer"
	order3 "gitlab.ozon.dev/ipogiba/homework/loms/internal/repository/db/order"
	stock3 "gitlab.ozon.dev/ipogiba/homework/loms/internal/repository/db/stock"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/service/order"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/service/stock"
	desc "gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			mw.Validate,
		),
	)

	reflection.Register(grpcServer)

	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, dbConnStr)
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}

	err = dbConn.Ping(ctx)
	if err != nil {
		log.Fatalln("Failed to ping:", err)
	}
	producer, err := producer.NewSyncProducer()
	if err != nil {
		log.Fatalln("Failed to create sync producer:", err)
	}

	orderRepo := order3.NewOrderRepository(dbConn)
	stockRepo := stock3.NewStockRepository(dbConn)

	orderService := order.NewOrderService(
		orderRepo,
		stockRepo,
		producer,
		topic,
	)
	stockService := stock.NewStockService(stockRepo)

	controller := loms.NewService(orderService, stockService)

	desc.RegisterLomsServer(grpcServer, controller)

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// создаем клиента на наш grpc сервер
	conn, err := grpc.NewClient(fmt.Sprintf(":%d", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}

	gwmux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(headerMatcher))

	if err = desc.RegisterLomsHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: mw.WithHTTPLoggingMiddleware(gwmux),
	}

	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr)

	log.Fatalln(gwServer.ListenAndServe())
}
