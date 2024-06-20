package loms

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	servicepb "gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
)

var _ servicepb.LomsServer = (*Service)(nil)

type OrderService interface {
	CreateOrder(ctx context.Context, userID int64, items []*model.Item) (int64, error)
	OrderInfo(ctx context.Context, orderID int64) (*model.Order, error)
	OrderPay(ctx context.Context, id int64) error
	OrderCancel(ctx context.Context, id int64) error
}

type StocksService interface {
	StockInfo(ctx context.Context, sku model.SKU) (uint64, error)
}

type Service struct {
	servicepb.UnimplementedLomsServer
	orderImpl  OrderService
	stocksImpl StocksService
}

func NewService(orderImpl OrderService, stocksImpl StocksService) *Service {
	return &Service{
		orderImpl:  orderImpl,
		stocksImpl: stocksImpl,
	}
}
