package order

import (
	"context"
	"github.com/pkg/errors"
	errorapp "route256/loms/internal/errors"
	"route256/loms/internal/model"
)

func (c *Order) OrderInfo(ctx context.Context, orderID int64) (*model.Order, error) {
	order, err := c.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, errorapp.ErrOrderNotFound) {
			return nil, errorapp.ErrOrderNotFound
		}
		return nil, errors.Wrap(err, "repository.GetOrderByID")
	}

	return order, nil
}
