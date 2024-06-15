package order

import (
	"context"
	"route256/loms/internal/model"
)

func (r *OrderInMemoryRepo) CreateOrder(ctx context.Context, userID int64, items []*model.Item) *model.Order {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	order := &model.Order{
		ID:     r.nextID,
		UserID: userID,
		Items:  items,
		Status: string(model.New),
	}

	r.orders[r.nextID] = order
	r.nextID++

	return order
}
