package cart

import (
	"context"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
)

func (r *InMemoryRepository) DeleteItemsByUserID(ctx context.Context, userID int64) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	_, ok := r.items[userID]
	if !ok {
		return errorapp.ErrNotFoundUser
	}

	delete(r.items, userID)

	return nil
}
