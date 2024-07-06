package stock

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *StockRepo) ReserveRemove(ctx context.Context, items []*model.Item) error {
	query := `UPDATE stocks SET
                  reserved = reserved - $1,
                  total_count = total_count - $1
              	  WHERE sku = $2`
	for _, item := range items {
		_, err := r.conn.Exec(ctx,
			query, item.Count, item.Count)
		if err != nil {
			return errors.Wrapf(err, "failed to reserve stock, sku: %d", item.SKU)
		}
	}

	return nil
}
