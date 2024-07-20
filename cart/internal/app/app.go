package app

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/http_handlers"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/middleware"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/client/loms"
	client "gitlab.ozon.dev/ipogiba/homework/cart/internal/client/product_service"
	repository "gitlab.ozon.dev/ipogiba/homework/cart/internal/repository/cart"
	service "gitlab.ozon.dev/ipogiba/homework/cart/internal/service/cart"
)

type App struct {
	mux        *http.ServeMux
	config     config
	server     *http.Server
	cartServer *http_handlers.Implementation
}

func NewApp(config config) (*App, error) {
	reviewsRepository := repository.NewRepository()
	productService, err := client.NewProductServiceClient(config.productAddr, config.productToken, config.getProductRPSLimit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize product service")
	}
	lomsService, err := loms.NewLomsServiceClient(config.configLomsService.lomsAddr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize loms service")
	}

	cartService := service.NewService(reviewsRepository, productService, lomsService)
	cartServer := http_handlers.NewCart(cartService)

	mux := http.NewServeMux()
	return &App{
		mux:        mux,
		config:     config,
		server:     &http.Server{Addr: config.addr, Handler: InitMiddlewares(mux)},
		cartServer: cartServer,
	}, nil
}

func (a *App) ListenAndServe() error {
	a.mux.HandleFunc(a.config.path.addItemToCart, a.cartServer.AddItemToCart)
	a.mux.HandleFunc(a.config.path.deleteProductFromCart, a.cartServer.DeleteProductFromCart)
	a.mux.HandleFunc(a.config.path.clearCart, a.cartServer.ClearCart)
	a.mux.HandleFunc(a.config.path.listCartProducts, a.cartServer.ListCartProducts)
	a.mux.HandleFunc(a.config.path.checkout, a.cartServer.Checkout)

	return a.server.ListenAndServe()
}

func (a *App) Close() error {
	return a.server.Close()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

func InitMiddlewares(router *http.ServeMux) http.Handler {
	return middleware.LoggingRequestMiddleware(router)
}
