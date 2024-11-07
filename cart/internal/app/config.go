package app

import (
	"fmt"

	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/definitions"
)

type (
	Options struct {
		Addr, ProductToken, ProductAddr, GrpcAddr string
	}

	configProductService struct {
		productToken, productAddr string
		getProductRPSLimit        int
		cacheSize                 uint
	}
	configLomsService struct {
		lomsAddr string
	}
	path struct {
		addItemToCart, deleteProductFromCart, clearCart, listCartProducts, checkout string
	}
	config struct {
		addr string
		configProductService
		configLomsService
		path path
	}
)

func NewConfig(opts Options) config {
	return config{
		addr: opts.Addr,
		configProductService: configProductService{
			productToken:       opts.ProductToken,
			productAddr:        opts.ProductAddr,
			getProductRPSLimit: 10,
			cacheSize:          100,
		},
		configLomsService: configLomsService{
			lomsAddr: opts.GrpcAddr,
		},
		path: path{
			addItemToCart:         fmt.Sprintf("POST /user/{%s}/cart/{%s}", definitions.ParamUserID, definitions.ParamSkuID),
			deleteProductFromCart: fmt.Sprintf("DELETE /user/{%s}/cart/{%s}", definitions.ParamUserID, definitions.ParamSkuID),
			clearCart:             fmt.Sprintf("DELETE /user/{%s}/cart", definitions.ParamUserID),
			listCartProducts:      fmt.Sprintf("GET /user/{%s}/cart/list", definitions.ParamUserID),
			checkout:              fmt.Sprintf("POST /user/{%s}/cart/checkout", definitions.ParamUserID),
		},
	}
}
