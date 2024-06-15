package order

import (
	"context"
	"fmt"
	errorapp "route256/loms/internal/errors"
	"route256/loms/internal/model"
)

func (r *OrderInMemoryRepo) SetStatus(ctx context.Context, id int64, status model.Status) error {
	order, ok := r.orders[id]
	if !ok {
		return errorapp.ErrOrderNotFound
	}

	switch status {
	case model.New, model.AwaitingPayment, model.Failed, model.Payed, model.Cancelled:
		order.Status = status.String()
		r.orders[id] = order
		return nil
	default:
		return fmt.Errorf("unknown status: %s", status)
	}
}
