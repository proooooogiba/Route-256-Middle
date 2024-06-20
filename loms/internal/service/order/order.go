package order

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

type IOrderRepo interface {
	CreateOrder(ctx context.Context, userID int64, items []*model.Item) *model.Order
	GetOrderByID(ctx context.Context, id int64) (*model.Order, error)
	SetStatus(ctx context.Context, id int64, status model.Status) error
}

type IStockRepo interface {
	Reserve(ctx context.Context, items []*model.Item) error
	ReserveRemove(ctx context.Context, items []*model.Item) error
	ReserveCancel(ctx context.Context, items []*model.Item) error
}

type Order struct {
	orderRepo IOrderRepo
	stockRepo IStockRepo
}

func NewOrderService(orderRepo IOrderRepo, stockRepo IStockRepo) *Order {
	return &Order{
		orderRepo: orderRepo,
		stockRepo: stockRepo,
	}
}
