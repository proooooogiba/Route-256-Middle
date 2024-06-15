package order

import (
	"context"
	"github.com/pkg/errors"
	"route256/loms/internal/model"
)

func (c *Order) OrderCancel(ctx context.Context, id int64) error {
	order, err := c.orderRepo.GetOrderByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "orderRepo.GetOrderByID")
	}

	err = c.stockRepo.ReserveCancel(ctx, order.Items)
	if err != nil {
		return errors.Wrap(err, "stockService.ReserveRemove")
	}

	err = c.orderRepo.SetStatus(ctx, order.ID, model.Cancelled)
	if err != nil {
		return errors.Wrap(err, "order.SetStatus")
	}

	return nil
}
