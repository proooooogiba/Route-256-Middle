package cart

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/app/http_handlers"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

func (s *cartServiceTestSuite) TestListProducts() {
	s.T().Run("failed on repo.GetItemsByUserID", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		s.cartRepository.GetItemsByUserIDMock.Expect(ctx, userID).Return(nil, ErrNoNil)
		product, err := s.srv.ListProducts(ctx, userID)
		assert.Error(t, err)
		assert.Nil(t, product)
	})

	s.T().Run("failed on productService.GetProductBySKU", func(t *testing.T) {
		var (
			ctx                 = context.Background()
			userID        int64 = 1
			returnedItems       = []model.Item{
				{
					SKU:   1,
					Count: 1,
				},
				{
					SKU:   2,
					Count: 2,
				},
			}
		)

		s.cartRepository.GetItemsByUserIDMock.Expect(ctx, userID).Return(returnedItems, nil)
		s.productService.GetProductBySKUMock.When(minimock.AnyContext, returnedItems[0].SKU).Then(nil, ErrNoNil)
		s.productService.GetProductBySKUMock.When(minimock.AnyContext, returnedItems[1].SKU).Then(nil, ErrNoNil)
		product, err := s.srv.ListProducts(ctx, userID)
		assert.Error(t, err)
		assert.Nil(t, product)
	})

	s.T().Run("successful", func(t *testing.T) {
		var (
			ctx                 = context.Background()
			userID        int64 = 1
			returnedItems       = []model.Item{
				{
					SKU:   1,
					Count: 1,
				},
				{
					SKU:   2,
					Count: 2,
				},
			}
			returnedProducts = []*model.Product{
				{
					SKU:   1,
					Name:  "11",
					Price: 111,
				},
				{
					SKU:   2,
					Name:  "22",
					Price: 222,
				},
			}
			expectedProducts = &http_handlers.ListCartProductsResponse{
				Items: []*http_handlers.Item{
					{
						SKU:   1,
						Name:  "11",
						Count: 1,
						Price: 111,
					},
					{
						SKU:   2,
						Name:  "22",
						Count: 2,
						Price: 222,
					},
				},
				TotalPrice: 555,
			}
		)

		s.cartRepository.GetItemsByUserIDMock.Expect(ctx, userID).Return(returnedItems, nil)
		s.productService.GetProductBySKUMock.When(minimock.AnyContext, returnedItems[0].SKU).Then(returnedProducts[0], nil)
		s.productService.GetProductBySKUMock.When(minimock.AnyContext, returnedItems[1].SKU).Then(returnedProducts[1], nil)

		products, err := s.srv.ListProducts(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, products)
	})
}
