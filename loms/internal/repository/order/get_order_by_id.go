package order

import (
	"context"
	errorapp "route256/loms/internal/errors"
	"route256/loms/internal/model"
)

func (r *OrderInMemoryRepo) GetOrderByID(ctx context.Context, id int64) (*model.Order, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	order, ok := r.orders[id]
	if !ok {
		return nil, errorapp.ErrOrderNotFound
	}

	return order, nil
}
