package cart

import (
	"context"
	errorapp "route256/cart/internal/errors"
	"route256/cart/internal/model"
)

func (r *InMemoryRepository) GetItemsByUserID(ctx context.Context, userID int64) ([]model.Item, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	cart, ok := r.items[userID]
	if !ok {
		return nil, errorapp.ErrNotFoundUser
	}
	return cart.Items, nil
}
