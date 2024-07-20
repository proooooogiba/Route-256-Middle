package stock

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (s *Stock) StockInfo(ctx context.Context, sku model.SKU) (uint64, error) {
	stock, err := s.stockRepo.GetStockBySKU(ctx, sku)
	if err != nil {
		return 0, errors.Wrap(err, "stockRepo.GetBySKU")
	}

	available := stock.TotalCount - stock.Reserved

	return available, nil
}
