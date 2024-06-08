package app

import (
	"net/http"
	"route256/cart/internal/app/http_handlers"
	"route256/cart/internal/app/middleware"
	client "route256/cart/internal/client/product_service"
	repository "route256/cart/internal/repository/cart"
	service "route256/cart/internal/service/cart"
)

type App struct {
	mux        *http.ServeMux
	config     config
	server     *http.Server
	cartServer *http_handlers.Implementation
}

func NewApp(config config) (*App, error) {
	reviewsRepository := repository.NewRepository()
	productService, err := client.NewProductServiceClient(config.productAddr, config.productToken)
	if err != nil {
		return nil, err
	}
	cartService := service.NewService(reviewsRepository, productService)
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

	return a.server.ListenAndServe()
}

func (a *App) Close() error {
	return a.server.Close()
}

func InitMiddlewares(router *http.ServeMux) http.Handler {
	return middleware.LoggingRequestMiddleware(router)
}
