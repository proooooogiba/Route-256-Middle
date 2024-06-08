package cart

import (
	"context"
	"route256/cart/internal/model"
)

type Repository interface {
	AddItem(ctx context.Context, userID int64, item model.Item) error
	DeleteItem(ctx context.Context, userID int64, sku model.SKU) error
	Clear(ctx context.Context, userID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) ([]model.Item, error)
}

type ProductService interface {
	GetProductBySKU(ctx context.Context, sku model.SKU) (*model.Product, error)
}

type Service struct {
	repo           Repository
	productService ProductService
}

func NewService(repository Repository, productService ProductService) *Service {
	return &Service{
		repo:           repository,
		productService: productService,
	}
}
