package order

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
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
	err = c.sendOrderEvent(ctx, order, model.AwaitingPayment)
	if err != nil {
		log.Printf("failed to send order event: %v", err)
	}

	return nil
}
