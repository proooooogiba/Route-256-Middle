package cart

import (
	"context"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
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
