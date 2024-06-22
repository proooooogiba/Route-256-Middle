package loms

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/cart/pkg/api/loms/v1"
)

func (c *LomsService) CreateOrder(ctx context.Context, userID int64, items []model.Item) (int64, error) {
	convertedItems := make([]*loms.Item, 0, len(items))
	for _, item := range items {
		convertedItems = append(convertedItems, &loms.Item{
			Sku:   int32(item.SKU),
			Count: uint32(item.Count),
		})
	}

	resp, err := c.client.CreateOrder(ctx, &loms.CreateOrderRequest{
		UserId: userID,
		Items:  convertedItems,
	})
	if err != nil {
		return 0, errors.Wrap(err, "client.CreateOrder")
	}

	return resp.GetOrderId(), nil
}
