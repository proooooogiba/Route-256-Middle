package main

import (
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"route256/cart/internal/app/cart/handlers"
	"route256/cart/internal/app/pkg/cart/client"
	"route256/cart/internal/app/pkg/cart/repository"
	"route256/cart/internal/app/pkg/cart/service"
	"route256/cart/internal/app/pkg/middleware"
)

func main() {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("error in logger start")
		return
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in logger sync")
		}
	}()
	logger.Info("starting app")

	conn, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		logger.Fatalw("error in listen", "err", err)
	}

	defer conn.Close()

	logger.Info("starting listen")

	reviewsRepository := repository.NewRepository()
	productService := client.NewProductServiceClient()
	cartService := service.NewService(reviewsRepository, productService)

	cartServer := handlers.NewCart(cartService)

	router := http.NewServeMux()
	router.HandleFunc("POST /user/{user_id}/cart/{sku_id}", cartServer.AddItemToCart)
	router.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", cartServer.DeleteProductFromCart)
	router.HandleFunc("DELETE /user/{user_id}/cart", cartServer.ClearCart)
	router.HandleFunc("GET /user/{user_id}/cart/list", cartServer.ListCartProducts)

	mux := middleware.LoggingRequestMiddleware(router)

	logger.Infow("starting server", "type", "START")

	if err := http.Serve(conn, mux); err != nil {
		logger.Fatalw("error in serve", "err", err)
	}
	return
}
