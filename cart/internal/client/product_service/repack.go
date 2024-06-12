package product_service

import "route256/cart/internal/model"

func getProduct(sku model.SKU, productRaw GetProductsResponse) *model.Product {
	return &model.Product{
		SKU:   sku,
		Name:  productRaw.Name,
		Price: productRaw.Price,
	}
}
