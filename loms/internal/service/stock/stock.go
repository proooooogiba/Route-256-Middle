package stock

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

type IStockRepo interface {
	GetStockBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error)
}

type Stock struct {
	stockRepo IStockRepo
}

func NewStockService(stockRepo IStockRepo) *Stock {
	return &Stock{
		stockRepo: stockRepo,
	}
}
