package product_service

import (
	"github.com/pkg/errors"
	"go.uber.org/ratelimit"
)

type ProductService struct {
	basePath          string
	token             string
	limiterGetProduct ratelimit.Limiter
}

func NewProductServiceClient(basePath string, token string, getProductRPSLimit int) (*ProductService, error) {
	if token == "" {
		return nil, errors.New("product service has empty auth token")
	}

	return &ProductService{
		token:             token,
		basePath:          basePath,
		limiterGetProduct: ratelimit.New(getProductRPSLimit),
	}, nil
}

type GetProductsRequest struct {
	Token string `json:"token"`
	Sku   uint32 `json:"sku"`
}

type GetProductsResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}
