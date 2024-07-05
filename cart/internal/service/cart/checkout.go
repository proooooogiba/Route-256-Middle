package cart

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/http_handlers"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
)

func (c *Service) Checkout(ctx context.Context, userID int64) (*http_handlers.CheckoutResponse, error) {
	items, err := c.repo.GetItemsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			return nil, err
		}
		return nil, errors.Wrap(err, "repo.GetItemsByUserID")
	}

	orderID, err := c.lomsService.CreateOrder(ctx, userID, items)
	if err != nil {
		return nil, errors.Wrap(err, "lomsService.CreateOrder")
	}

	err = c.repo.DeleteItemsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundUser) {
			return nil, err
		}
		return nil, err
	}

	return &http_handlers.CheckoutResponse{
		OrderID: orderID,
	}, nil
}
