package order

import (
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"sync"
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
