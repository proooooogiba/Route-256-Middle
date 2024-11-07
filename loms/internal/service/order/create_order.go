package order

import (
	"context"
	"log"

	"github.com/pkg/errors"
	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Order) CreateOrder(ctx context.Context, userID int64, items []*model.Item) (int64, error) {
	order, err := c.orderRepo.CreateOrder(ctx, userID, items)
	if err != nil {
		return 0, errors.Wrapf(err, "orderRepo.CreateOrder")
	}
	err = c.sendOrderEvent(ctx, order, model.New)
	if err != nil {
		log.Printf("failed to send order event: %v", err)
	}

	err = c.stockRepo.Reserve(ctx, order.Items)
	if err != nil {
		errStatus := c.orderRepo.SetStatus(ctx, order.ID, model.Failed)
		if errStatus != nil {
			return 0, errors.Wrap(errStatus, "orderRepo.SetStatus")
		}
		err = c.sendOrderEvent(ctx, order, model.Failed)
		if err != nil {
			log.Printf("failed to send order event: %v", err)
		}

		if errors.Is(err, errorapp.ErrSkuNotFound) || errors.Is(err, errorapp.ErrOutOfStock) {
			return 0, err
		}
		return 0, errors.Wrap(err, "stockRepo.Reserve")
	}

	err = c.orderRepo.SetStatus(ctx, order.ID, model.AwaitingPayment)
	if err != nil {
		return 0, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	err = c.sendOrderEvent(ctx, order, model.AwaitingPayment)
	if err != nil {
		log.Printf("failed to send order event: %v", err)
	}

	return order.ID, nil
}
