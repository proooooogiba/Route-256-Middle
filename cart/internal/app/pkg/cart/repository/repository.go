package repository

import (
	"context"
	errorapp "route256/cart/internal/app/pkg/errors"
	"route256/cart/internal/app/pkg/model"
	"sort"
)

type CartStorage map[int64]model.Cart

type InMemoryRepository struct {
	storage CartStorage
}

func NewRepository() *InMemoryRepository {
	return &InMemoryRepository{
		storage: make(CartStorage, 10),
	}
}

func (r *InMemoryRepository) AddItem(ctx context.Context, userID int64, item model.Item) error {
	cart, ok := r.storage[userID]
	if !ok {
		r.storage[userID] = model.Cart{
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

	r.storage[userID] = cart

	return nil
}

func (r *InMemoryRepository) DeleteItem(ctx context.Context, userID int64, sku model.SKU) error {
	cart, ok := r.storage[userID]
	if !ok {
		return errorapp.ErrNotFoundUser
	}
	for _, cartItem := range cart.Items {
		if cartItem.SKU == sku {
			delete(r.storage, userID)
			return nil
		}
	}

	return nil
}

func (r *InMemoryRepository) Clear(ctx context.Context, userID int64) error {
	_, ok := r.storage[userID]
	if !ok {
		return nil
	}

	delete(r.storage, userID)

	return nil
}

func (r *InMemoryRepository) GetItemsByUserID(ctx context.Context, userID int64) ([]model.Item, error) {
	cart, ok := r.storage[userID]
	if !ok {
		return nil, errorapp.ErrNotFoundUser
	}
	return cart.Items, nil
}
