package cart

import (
	"context"
	"sort"

	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

func (r *InMemoryRepository) AddItem(ctx context.Context, userID int64, item model.Item) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	cart, ok := r.items[userID]
	if !ok {
		r.items[userID] = model.Cart{
			Items: []model.Item{item},
		}
		return nil
	}

	found := false
	for i := range cart.Items {
		if cart.Items[i].SKU == item.SKU {
			cart.Items[i].Count += item.Count
			found = true
			break
		}
	}
	if !found {
		cart.Items = append(cart.Items, item)
	}

	sort.Slice(cart.Items, func(i, j int) bool {
		return cart.Items[i].SKU < cart.Items[j].SKU
	})

	r.items[userID] = cart

	return nil
}
