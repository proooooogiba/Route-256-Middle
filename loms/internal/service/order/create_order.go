package order

import (
	"context"
	"github.com/pkg/errors"
	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

func (c *Order) CreateOrder(ctx context.Context, userID int64, items []*model.Item) (int64, error) {
	order, err := c.orderRepo.CreateOrder(ctx, userID, items)
	if err != nil {
		return 0, errors.Wrapf(err, "orderRepo.CreateOrder")
	}

	err = c.stockRepo.Reserve(ctx, order.Items)
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
