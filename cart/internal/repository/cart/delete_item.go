package cart

import (
	"context"
	errorapp "route256/cart/internal/errors"
	"route256/cart/internal/model"
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
