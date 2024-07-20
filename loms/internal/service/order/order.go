package order

import (
	"context"

	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

type IOrderRepo interface {
	CreateOrder(ctx context.Context, userID int64, items []*model.Item) (*model.Order, error)
	GetOrderByID(ctx context.Context, id int64) (*model.Order, error)
	SetStatus(ctx context.Context, id int64, status model.Status) error
}

type IStockRepo interface {
	Reserve(ctx context.Context, items []*model.Item) error
	ReserveRemove(ctx context.Context, items []*model.Item) error
	ReserveCancel(ctx context.Context, items []*model.Item) error
}

type IProducer interface {
	SendMessages(ctx context.Context, msgs []model.ProducerMessage) error
	SendMessage(ctx context.Context, msg model.ProducerMessage) error
}

type Order struct {
	orderRepo IOrderRepo
	stockRepo IStockRepo
	producer  IProducer
	topic     string
}

func NewOrderService(
	orderRepo IOrderRepo,
	stockRepo IStockRepo,
	producer IProducer,
	topic string,
) *Order {
	return &Order{
		orderRepo: orderRepo,
		stockRepo: stockRepo,
		producer:  producer,
		topic:     topic,
	}
}
