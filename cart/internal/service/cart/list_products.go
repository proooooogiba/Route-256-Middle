package cart

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/http_handlers"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
)

func (c *Service) ListProducts(ctx context.Context, userID int64) (*http_handlers.ListCartProductsResponse, error) {
	items, err := c.repo.GetItemsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			return nil, err
		}
		return nil, errors.Wrap(err, "repo.GetItemsByUserID")
	}

	var totalPrice uint32
	respItems := make([]*http_handlers.Item, len(items))
	for i, item := range items {
		product, err := c.productService.GetProductBySKU(ctx, item.SKU)
		if err != nil {
			return nil, errors.Wrap(err, "productService.GetProductBySKU")
		}
		respItems[i] = &http_handlers.Item{
			SKU:   product.SKU,
			Name:  product.Name,
			Count: item.Count,
			Price: product.Price,
		}
		totalPrice += product.Price * uint32(item.Count)
	}

	resp := &http_handlers.ListCartProductsResponse{
		Items:      respItems,
		TotalPrice: totalPrice,
	}

	return resp, nil
}
