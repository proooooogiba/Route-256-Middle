package product_service_in_memory_cache

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/cache"
	client "gitlab.ozon.dev/ipogiba/homework/cart/internal/client/product_service"
)

type ProductServiceInMemoryCache struct {
	productService *client.ProductService
	cache          cache.Cacher
}

func NewProductServiceInMemoryCacheClient(basePath string, token string, getProductRPSLimit int, cache cache.Cacher) (*ProductServiceInMemoryCache, error) {
	if token == "" {
		return nil, errors.New("product service has empty auth token")
	}

	productService, err := client.NewProductServiceClient(basePath, token, getProductRPSLimit)
	if err != nil {
		return nil, errors.Wrap(err, "new product service client")
	}

	return &ProductServiceInMemoryCache{
		productService: productService,
		cache:          cache,
	}, nil
}
