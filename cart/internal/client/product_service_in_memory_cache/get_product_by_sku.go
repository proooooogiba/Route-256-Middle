package product_service_in_memory_cache

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"log"
)

func (c *ProductServiceInMemoryCache) GetProductBySKU(ctx context.Context, sku model.SKU) (*model.Product, error) {
	var (
		product *model.Product
		err     error
	)

	product, err = c.cache.Get(ctx, sku)
	if err == nil {
		return product, nil
	}

	product, err = c.productService.GetProductBySKU(ctx, sku)
	if err != nil {
		return nil, err
	}
	c.saveToCache(ctx, product)

	return product, nil
}

func (c *ProductServiceInMemoryCache) saveToCache(ctx context.Context, product *model.Product) {
	go func() {
		err := c.cache.Set(ctx, product)
		if err != nil {
			log.Fatalf("saveToCache error: %+v", err)
		}
	}()
}
