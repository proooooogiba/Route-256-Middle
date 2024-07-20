package order

import (
	"context"

	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
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
