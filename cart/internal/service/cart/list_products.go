package cart

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/http_handlers"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/pkg/errgroup"
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

	mu := sync.Mutex{}
	group := errgroup.New(ctx)
	for i, item := range items {
		group.Go(func() error {
			product, err := c.productService.GetProductBySKU(ctx, item.SKU)
			if err != nil {
				return errors.Wrap(err, "productService.GetProductBySKU")
			}

			respItem := &http_handlers.Item{
				SKU:   product.SKU,
				Name:  product.Name,
				Count: item.Count,
				Price: product.Price,
			}

			mu.Lock()
			defer mu.Unlock()
			respItems[i] = respItem
			totalPrice += product.Price * uint32(item.Count)

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return nil, errors.Wrap(err, "group.Wait")
	}

	resp := &http_handlers.ListCartProductsResponse{
		Items:      respItems,
		TotalPrice: totalPrice,
	}

	return resp, nil
}
