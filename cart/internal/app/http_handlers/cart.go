package http_handlers

import (
	"context"
	"route256/cart/internal/model"
)

type CartServer interface {
	AddItem(ctx context.Context, userID int64, sku model.SKU, count uint16) error
	DeleteItem(ctx context.Context, userID int64, sku model.SKU) error
	Clear(ctx context.Context, userID int64) error
	ListProducts(ctx context.Context, userID int64) (*ListCartProductsResponse, error)
}

type Implementation struct {
	cartService CartServer
}

func NewCart(
	cartService CartServer,
) *Implementation {
	return &Implementation{cartService: cartService}
}
