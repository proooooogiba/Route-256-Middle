package stock

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *StockRepo) Reserve(ctx context.Context, items []*model.Item) error {
	for _, item := range items {
		_, err := r.conn.Exec(ctx, `UPDATE stocks SET reserved = reserved + $1 WHERE sku = $2`, item.Count, item.Count)
		if err != nil {
			return errors.Wrapf(err, "failed to reserve stock, sku: %d", item.SKU)
		}
	}

	return nil
}
