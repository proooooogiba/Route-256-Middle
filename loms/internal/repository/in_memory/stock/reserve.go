package stock

import (
	"context"

	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *StockInMemoryRepo) Reserve(ctx context.Context, items []*model.Item) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	tempStock := make(map[model.SKU]*model.Stock, len(items))
	for sku, stock := range r.stocks {
		tempStock[sku] = stock
	}

	for _, item := range items {
		stock, ok := tempStock[item.SKU]
		if !ok {
			return errorapp.ErrSkuNotFound
		}

		if stock.TotalCount-stock.Reserved < uint64(item.Count) {
			return errorapp.ErrOutOfStock
		}

		stock.Reserved += uint64(item.Count)
		tempStock[item.SKU] = stock
	}

	r.stocks = tempStock

	return nil
}
