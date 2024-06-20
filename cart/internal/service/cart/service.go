package cart

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

type Repository interface {
	AddItem(ctx context.Context, userID int64, item model.Item) error
	DeleteItem(ctx context.Context, userID int64, sku model.SKU) error
	DeleteItemsByUserID(ctx context.Context, userID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) ([]model.Item, error)
}

type ProductService interface {
	GetProductBySKU(ctx context.Context, sku model.SKU) (*model.Product, error)
}

type LomsService interface {
	CreateOrder(ctx context.Context, userID int64, item []model.Item) error
	StocksInfo(ctx context.Context, sku model.SKU) (model.StockItem, error)
}

type Service struct {
	repo           Repository
	productService ProductService
	lomsService    LomsService
}

func NewService(repository Repository, productService ProductService, lomsService LomsService) *Service {
	return &Service{
		repo:           repository,
		productService: productService,
		lomsService:    lomsService,
	}
}
