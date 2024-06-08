package main

import (
	"flag"
	"fmt"
	"route256/cart/internal/app"
)

const (
	defaultAddr        = ":8082"
	defaultProductAddr = "http://route256.pavl.uk:8080/"

	envToken = "TOKEN"
	token    = "testtoken"
)

var opts = app.Options{}

func initOpts() {
	flag.StringVar(&opts.Addr, "addr", defaultAddr, fmt.Sprintf("server address, default: %q", defaultAddr))
	flag.StringVar(&opts.ProductAddr, "product_addr", defaultProductAddr, fmt.Sprintf("products-service address, default: %q", defaultProductAddr))
	flag.Parse()

	opts.ProductToken = token
}
