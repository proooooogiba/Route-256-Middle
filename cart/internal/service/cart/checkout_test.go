package cart

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
)

func (s *cartServiceTestSuite) TestCheckout() {
	s.T().Run("failed on repo.GetItemsByUserID", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		s.cartRepository.GetItemsByUserIDMock.Expect(ctx, userID).Return(nil, ErrNoNil)
		resp, err := s.srv.Checkout(ctx, userID)

		//require.Equal(t, resp.OrderID, 1)
		require.Nil(t, resp)
		require.Error(t, err)
	})

	s.T().Run("failed on lomsService.CreateOrder", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
			items        = []model.Item{
				{
					SKU:   1,
					Count: 1,
				},
			}
		)

		s.cartRepository.GetItemsByUserIDMock.Expect(ctx, userID).Return(items, nil)
		s.lomsService.CreateOrderMock.Expect(ctx, userID, items).Return(0, ErrNoNil)

		resp, err := s.srv.Checkout(ctx, userID)
		require.Nil(t, resp)
		require.Error(t, err)
	})

	s.T().Run("failed on repo.DeleteItemsByUserID", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
			items        = []model.Item{
				{
					SKU:   1,
					Count: 1,
				},
			}
		)

		s.cartRepository.GetItemsByUserIDMock.Expect(ctx, userID).Return(items, nil)
		s.lomsService.CreateOrderMock.Expect(ctx, userID, items).Return(0, nil)
		s.cartRepository.DeleteItemsByUserIDMock.Expect(ctx, userID).Return(ErrNoNil)

		resp, err := s.srv.Checkout(ctx, userID)
		//require.Equal(t, resp.OrderID, 1)
		require.Nil(t, resp)
		require.Error(t, err)
	})

	s.T().Run("successful", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		s.cartRepository.DeleteItemsByUserIDMock.Expect(ctx, userID).Return(nil)

		err := s.srv.Clear(ctx, userID)
		require.NoError(t, err)
	})
	s.T().Run("successful", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
			items        = []model.Item{
				{
					SKU:   1,
					Count: 1,
				},
			}
		)

		s.cartRepository.GetItemsByUserIDMock.Expect(ctx, userID).Return(items, nil)
		s.lomsService.CreateOrderMock.Expect(ctx, userID, items).Return(1, nil)
		s.cartRepository.DeleteItemsByUserIDMock.Expect(ctx, userID).Return(nil)

		resp, err := s.srv.Checkout(ctx, userID)
		require.EqualValues(t, 1, resp.OrderID)
		require.NoError(t, err)
	})
}
