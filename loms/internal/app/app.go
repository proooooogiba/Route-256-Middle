package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/pkg/shard_manager"
	order3 "gitlab.ozon.dev/ipogiba/homework/loms/internal/repository/db/order"
	stock2 "gitlab.ozon.dev/ipogiba/homework/loms/internal/repository/db/stock"

	"github.com/ds248a/closer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/app/loms"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/mw"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/producer"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/service/order"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/service/stock"
	desc "gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	config     Config
	grpcServer *grpc.Server
	lis        net.Listener

	gatewayServer *http.Server
}

func New(ctx context.Context, config Config) (*App, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			mw.Validate,
		),
	)

	reflection.Register(grpcServer)

	databases, err := initDBPool(ctx, config.databasePool)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init db pool")
	}

	sm := shard_manager.New(
		shard_manager.GetMurmur3ShardFn(len(databases)),
		databases,
	)

	orderRepo := order3.NewOrderRepository(sm)
	stockRepo := stock2.NewStockRepository(databases[0])

	producer, err := producer.NewSyncProducer(config.brokers)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sync producer")
	}
	closer.Add(producer.Close)

	orderService := order.NewOrderService(
		orderRepo,
		stockRepo,
		producer,
		config.topic,
	)
	stockService := stock.NewStockService(stockRepo)

	controller := loms.NewService(orderService, stockService)

	desc.RegisterLomsServer(grpcServer, controller)

	log.Printf("server listening at %v", lis.Addr())

	return &App{
		config:     config,
		grpcServer: grpcServer,
		lis:        lis,
	}, nil
}

func (a *App) Close() {
	closer.Close(syscall.SIGTERM)
}

func (a *App) ServeGrpcServer() {
	go func() {
		if err := a.grpcServer.Serve(a.lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

// создаем клиента на наш grpc сервер
func (a *App) CreateGatewayClient(ctx context.Context) error {
	conn, err := grpc.NewClient(
		fmt.Sprintf(":%s", a.config.grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create gRPC client")
	}

	gwmux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(headerMatcher))

	if err = desc.RegisterLomsHandler(ctx, gwmux, conn); err != nil {
		return errors.Wrap(err, "failed to register gateway")
	}

	a.gatewayServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.httpPort),
		Handler: InitMiddleware(gwmux),
	}

	return nil
}

func (a *App) GatewayListenAndServe() error {
	return a.gatewayServer.ListenAndServe()
}

func InitMiddleware(mux http.Handler) http.Handler {
	return mw.WithHTTPLoggingMiddleware(mux)
}

func initDBConnect(ctx context.Context, dbConnStr string) (*pgx.Conn, error) {
	dbConn, err := pgx.Connect(ctx, dbConnStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to databasePool")
	}

	err = dbConn.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping databasePool")
	}

	return dbConn, nil
}

func initDBPool(ctx context.Context, databasePool []database) ([]*pgxpool.Pool, error) {
	if len(databasePool) == 0 {
		return nil, errors.New("no database pool provided")
	}

	var databases []*pgxpool.Pool
	for _, dbConf := range databasePool {
		db, err := pgxpool.New(ctx, dbConf.DSN)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to connect to database %q", dbConf.DSN)
		}
		closer.Add(db.Close)

		databases = append(databases, db)
	}

	return databases, nil
}

func headerMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case "x-auth":
		return key, true
	default:
		return key, false
	}
}
