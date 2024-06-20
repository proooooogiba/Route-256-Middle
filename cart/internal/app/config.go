package app

import (
	"fmt"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/definitions"
)

type (
	Options struct {
		Addr, ProductToken, ProductAddr string
	}

	configProductService struct {
		productToken, productAddr string
	}
	path struct {
		addItemToCart, deleteProductFromCart, clearCart, listCartProducts string
	}
	config struct {
		addr string
		configProductService
		path path
	}
)

func NewConfig(opts Options) config {
	return config{
		addr: opts.Addr,
		configProductService: configProductService{
			productToken: opts.ProductToken,
			productAddr:  opts.ProductAddr,
		},
		path: path{
			addItemToCart:         fmt.Sprintf("POST /user/{%s}/cart/{%s}", definitions.ParamUserID, definitions.ParamSkuID),
			deleteProductFromCart: fmt.Sprintf("DELETE /user/{%s}/cart/{%s}", definitions.ParamUserID, definitions.ParamSkuID),
			clearCart:             fmt.Sprintf("DELETE /user/{%s}/cart", definitions.ParamUserID),
			listCartProducts:      fmt.Sprintf("GET /user/{%s}/cart/list", definitions.ParamUserID),
		},
	}
}
