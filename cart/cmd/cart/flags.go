package main

import (
	"flag"
	"fmt"

	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app"
)

const (
	defaultAddr        = ":8082"
	defaultProductAddr = "http://route256.pavl.uk:8080/"
	grpcLomsAddr       = "loms:50051"

	envToken = "TOKEN"
	token    = "testtoken"
)

var opts = app.Options{}

func initOpts() {
	flag.StringVar(&opts.Addr, "addr", defaultAddr, fmt.Sprintf("server address, default: %q", defaultAddr))
	flag.StringVar(&opts.GrpcAddr, "grpc-addr", grpcLomsAddr, fmt.Sprintf("grpc address: %q", grpcLomsAddr))
	flag.StringVar(&opts.ProductAddr, "product_addr", defaultProductAddr, fmt.Sprintf("products-service address, default: %q", defaultProductAddr))
	flag.Parse()

	opts.ProductToken = token
}
