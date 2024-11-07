package cart

import (
	"context"

	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

func (r *InMemoryRepository) DeleteItem(ctx context.Context, userID int64, sku model.SKU) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	cart, ok := r.items[userID]
	if !ok {
		return errorapp.ErrNotFoundUser
	}

	for i, cartItem := range cart.Items {
		if cartItem.SKU == sku {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			r.items[userID] = cart
			return nil
		}
	}

	return nil
}
