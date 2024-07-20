package stock

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *StockRepo) GetStockBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error) {
	var stock model.Stock
	err := r.conn.QueryRow(ctx, `SELECT sku, total, reserved FROM stocks WHERE sku = $1`, sku).
		Scan(
			&stock.Sku,
			&stock.TotalCount,
			&stock.Reserved,
		)
	if err != nil {
		return nil, errors.Wrap(err, "select stock")
	}

	return &stock, nil
}
