package service

import (
	"context"
	"github.com/pkg/errors"
	"route256/cart/internal/app/cart/handlers"
	"route256/cart/internal/app/pkg/model"
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

func (c *Service) AddItem(ctx context.Context, userID int64, sku model.SKU, count uint16) error {
	_, err := c.productService.GetProductBySKU(ctx, sku)
	if err != nil {
		return errors.Wrap(err, "productService.GetProductBySKU")
	}

	item := model.Item{
		SKU:   sku,
		Count: count,
	}

	err = c.repo.AddItem(ctx, userID, item)
	if err != nil {
		return errors.Wrap(err, "repo.AddItemToCart")
	}

	return nil
}

func (c *Service) DeleteItem(ctx context.Context, userID int64, sku model.SKU) error {
	return c.repo.DeleteItem(ctx, userID, sku)
}

func (c *Service) Clear(ctx context.Context, userID int64) error {
	return c.repo.Clear(ctx, userID)
}

func (c *Service) ListProducts(ctx context.Context, userID int64) (*handlers.ListCartProductsResponse, error) {
	items, err := c.repo.GetItemsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "repo.GetItemsByUserID")
	}

	var totalPrice uint32
	respItems := make([]*handlers.Item, len(items))
	for i, item := range items {
		product, err := c.productService.GetProductBySKU(ctx, item.SKU)
		if err != nil {
			return nil, errors.Wrap(err, "productService.GetProductBySKU")
		}
		respItems[i] = &handlers.Item{
			SKU:   product.SKU,
			Name:  product.Name,
			Count: item.Count,
			Price: product.Price,
		}
		totalPrice += product.Price
	}

	resp := &handlers.ListCartProductsResponse{
		Items:      respItems,
		TotalPrice: totalPrice,
	}

	return resp, nil
}
