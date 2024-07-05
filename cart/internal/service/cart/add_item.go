package cart

import (
	"context"
	"github.com/pkg/errors"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

func (c *Service) AddItem(ctx context.Context, userID int64, sku model.SKU, count uint16) error {
	_, err := c.productService.GetProductBySKU(ctx, sku)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundInPS) {
			return err
		}
		return errors.Wrap(err, "productService.GetProductBySKU")
	}

	stockItem, err := c.lomsService.StocksInfo(ctx, sku)
	if err != nil {
		return errors.Wrap(err, "lomsService.StocksInfo")
	}

	if stockItem.Count < uint64(count) {
		return errorapp.ErrOutOfStock
	}

	item := model.Item{
		SKU:   sku,
		Count: count,
	}

	err = c.repo.AddItem(ctx, userID, item)
	if err != nil {
		return errors.Wrap(err, "repo.AddItemToCart")
	}

	return nil
}
