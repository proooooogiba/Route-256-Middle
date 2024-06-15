package order

import (
	"route256/loms/internal/model"
	"sync"
)

type OrderInMemoryRepo struct {
	orders map[int64]*model.Order
	nextID int64
	mtx    sync.RWMutex
}

func NewStockRepository() (*OrderInMemoryRepo, error) {
	return &OrderInMemoryRepo{
		orders: make(map[int64]*model.Order),
		mtx:    sync.RWMutex{},
		nextID: 1,
	}, nil
}
