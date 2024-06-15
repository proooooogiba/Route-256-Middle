package stock

import (
	"context"
	errorapp "route256/loms/internal/errors"
	"route256/loms/internal/model"
)

func (r *StockInMemoryRepo) GetStockBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	stock, ok := r.stocks[sku]
	if !ok {
		return nil, errorapp.ErrStockNotFound
	}

	return stock, nil
}
