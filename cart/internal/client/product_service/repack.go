package product_service

import "gitlab.ozon.dev/ipogiba/homework/cart/internal/model"

func getProduct(sku model.SKU, productRaw GetProductsResponse) *model.Product {
	return &model.Product{
		SKU:   sku,
		Name:  productRaw.Name,
		Price: productRaw.Price,
	}
}
