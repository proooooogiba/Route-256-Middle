package order

import (
	"context"
	"github.com/pkg/errors"
	errorapp "route256/loms/internal/errors"
	"route256/loms/internal/model"
)

func (c *Order) CreateOrder(ctx context.Context, userID int64, items []*model.Item) (int64, error) {
	order := c.orderRepo.CreateOrder(ctx, userID, items)

	err := c.stockRepo.Reserve(ctx, order.Items)
	if err != nil {
		errStatus := c.orderRepo.SetStatus(ctx, order.ID, model.Failed)
		if errStatus != nil {
			return 0, errors.Wrap(errStatus, "orderRepo.SetStatus")
		}

		if errors.Is(err, errorapp.ErrSkuNotFound) || errors.Is(err, errorapp.ErrOutOfStock) {
			return 0, err
		}
		return 0, errors.Wrap(err, "stockRepo.Reserve")
	}

	err = c.orderRepo.SetStatus(ctx, order.ID, model.AwaitingPayment)
	if err != nil {
		return 0, errors.Wrap(err, "orderRepo.SetStatus")
	}

	return order.ID, nil
}
