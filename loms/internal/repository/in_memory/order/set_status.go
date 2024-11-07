package order

import (
	"context"
	"fmt"

	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
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
