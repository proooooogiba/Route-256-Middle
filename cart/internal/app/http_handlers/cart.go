package http_handlers

import (
	"context"

	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

type CartService interface {
	AddItem(ctx context.Context, userID int64, sku model.SKU, count uint16) error
	DeleteItem(ctx context.Context, userID int64, sku model.SKU) error
	Clear(ctx context.Context, userID int64) error
	ListProducts(ctx context.Context, userID int64) (*ListCartProductsResponse, error)
	Checkout(ctx context.Context, userID int64) (*CheckoutResponse, error)
}

type Implementation struct {
	cartService CartService
}

func NewCart(
	cartService CartService,
) *Implementation {
	return &Implementation{cartService: cartService}
}
