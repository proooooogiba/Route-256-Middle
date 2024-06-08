package cart

import (
	"context"
	"github.com/pkg/errors"
	errorapp "route256/cart/internal/errors"
	"route256/cart/internal/model"
)

func (c *Service) AddItem(ctx context.Context, userID int64, sku model.SKU, count uint16) error {
	_, err := c.productService.GetProductBySKU(ctx, sku)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundInPS) {
			return err
		}
		return errors.Wrap(err, "productService.GetProductBySKU")
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
