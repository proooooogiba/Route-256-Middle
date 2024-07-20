package cart

import (
	"sync"

	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

type InMemoryRepository struct {
	items map[int64]model.Cart
	mtx   sync.RWMutex
}

func NewRepository() *InMemoryRepository {
	return &InMemoryRepository{
		items: make(map[int64]model.Cart, 10),
		mtx:   sync.RWMutex{},
	}
}
