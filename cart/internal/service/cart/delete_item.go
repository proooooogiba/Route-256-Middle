package cart

import (
	"context"
	"route256/cart/internal/model"
)

func (c *Service) DeleteItem(ctx context.Context, userID int64, sku model.SKU) error {
	return c.repo.DeleteItem(ctx, userID, sku)
}
