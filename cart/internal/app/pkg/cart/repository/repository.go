package repository

import (
	"context"
	"errors"
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

func (r *InMemoryRepository) AddItem(ctx context.Context, userID int64, sku model.Item) error {
	// validation
	cart, ok := r.storage[userID]
	if !ok {
		r.storage[userID] = model.Cart{
			Items: []model.Item{sku},
		}
	}

	found := false
	for _, cartItem := range cart.Items {
		if cartItem.SKU == sku.SKU {
			cartItem.Count += sku.Count
			found = true
			break
		}
	}
	if !found {
		cart.Items = append(cart.Items, sku)
	}

	sort.Slice(cart.Items, func(i, j int) bool {
		return cart.Items[i].SKU < cart.Items[j].SKU
	})

	r.storage[userID] = cart

	return nil
}

func (r *InMemoryRepository) DeleteItem(ctx context.Context, userID int64, sku model.SKU) error {
	// validation
	cart, _ := r.storage[userID]
	for _, cartItem := range cart.Items {
		if cartItem.SKU == sku {
			delete(r.storage, userID)
			break
		}
	}

	r.storage[userID] = cart

	return nil
}

func (r *InMemoryRepository) Clear(ctx context.Context, userID int64) error {
	// validation
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
		return nil, errNotFound
	}
	return cart.Items, nil
}

var errNotFound = errors.New("not found")
