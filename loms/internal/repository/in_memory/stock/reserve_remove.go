package stock

import (
	"context"

	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (r *StockInMemoryRepo) ReserveRemove(ctx context.Context, items []*model.Item) error {
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

		stock.Reserved -= uint64(item.Count)
		stock.TotalCount -= uint64(item.Count)
		tempStock[item.SKU] = stock
	}

	r.stocks = tempStock

	return nil
}
