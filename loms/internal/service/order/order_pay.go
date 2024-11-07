package order

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (c *Order) OrderPay(ctx context.Context, id int64) error {
	order, err := c.orderRepo.GetOrderByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "orderRepo.GetOrderByID")
	}

	if order.Status != model.AwaitingPayment.String() {
		return errors.New("order status is not awaiting payment")
	}

	err = c.stockRepo.ReserveRemove(ctx, order.Items)
	if err != nil {
		return errors.Wrap(err, "stockService.ReserveRemove")
	}

	err = c.orderRepo.SetStatus(ctx, order.ID, model.Payed)
	if err != nil {
		return errors.Wrap(err, "order.SetStatus")
	}
	err = c.sendOrderEvent(ctx, order, model.Payed)
	if err != nil {
		log.Printf("failed to send order event: %v", err)
	}

	return nil
}
