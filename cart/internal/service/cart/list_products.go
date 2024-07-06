package cart

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/http_handlers"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/pkg/pipeline"
)

type productResult struct {
	product *model.Product
	err     error
}

func (c *Service) ListProducts(ctx context.Context, userID int64) (*http_handlers.ListCartProductsResponse, error) {
	items, err := c.repo.GetItemsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			return nil, err
		}
		return nil, errors.Wrap(err, "repo.GetItemsByUserID")
	}

	task := func(ctx context.Context, index int) (any, error) {
		item := items[index]
		product, err := c.productService.GetProductBySKU(ctx, item.SKU)
		return productResult{product: product, err: err}, err
	}

	var results []pipeline.Result
	results, err = pipeline.Parallelize(ctx, len(items), task)
	if err != nil {
		return nil, errors.Wrap(err, "pipeline.Parallelize")
	}

	var totalPrice uint32
	respItems := make([]*http_handlers.Item, len(items))
	for i, res := range results {
		if res.Err != nil {
			if errors.Is(res.Err, context.Canceled) {
				return nil, res.Err
			}
			return nil, errors.Wrap(res.Err, "task execution failed")
		}

		result := res.Value.(productResult)
		respItems[i] = &http_handlers.Item{
			SKU:   result.product.SKU,
			Name:  result.product.Name,
			Count: items[i].Count,
			Price: result.product.Price,
		}
		totalPrice += result.product.Price * uint32(items[i].Count)
	}

	resp := &http_handlers.ListCartProductsResponse{
		Items:      respItems,
		TotalPrice: totalPrice,
	}

	return resp, nil
}
