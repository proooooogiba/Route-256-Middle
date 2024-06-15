package order

import (
	"context"
	"github.com/pkg/errors"
	"route256/loms/internal/model"
)

func (c *Order) OrderPay(ctx context.Context, id int64) error {
	order, err := c.orderRepo.GetOrderByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "orderRepo.GetOrderByID")
	}

	payment := model.AwaitingPayment
	if order.Status != payment.String() {
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

	return nil
}
