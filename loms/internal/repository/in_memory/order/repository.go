package order

import (
	"sync"

	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
)

type OrderInMemoryRepo struct {
	orders map[int64]*model.Order
	nextID int64
	mtx    sync.RWMutex
}

func NewOrderRepository() *OrderInMemoryRepo {
	return &OrderInMemoryRepo{
		orders: make(map[int64]*model.Order),
		mtx:    sync.RWMutex{},
		nextID: 1,
	}
}
