package loms

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/cart/pkg/api/loms/v1"
)

func (c *LomsService) StocksInfo(ctx context.Context, sku model.SKU) (*model.StockItem, error) {
	resp, err := c.client.StocksInfo(ctx, &loms.StocksInfoRequest{
		Sku: int64(sku),
	})
	if err != nil {
		return nil, errors.Wrap(err, "client.StocksInfo")
	}

	return &model.StockItem{
		Sku:   sku,
		Count: resp.Count,
	}, nil
}
