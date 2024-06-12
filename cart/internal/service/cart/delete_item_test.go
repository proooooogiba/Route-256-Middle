package cart

import (
	"context"
	"github.com/stretchr/testify/require"
	"route256/cart/internal/model"
	"testing"
)

func (s *cartServiceTestSuite) TestDeleteItem() {
	s.T().Run("failed on repo.DeleteItem", func(t *testing.T) {
		var (
			ctx              = context.Background()
			userID int64     = 1
			sku    model.SKU = 1
		)

		s.cartRepository.DeleteItemMock.Expect(ctx, userID, sku).Return(ErrNoNil)
		err := s.srv.DeleteItem(ctx, userID, sku)
		require.Error(t, err)
	})

	s.T().Run("successful", func(t *testing.T) {
		var (
			ctx              = context.Background()
			userID int64     = 1
			sku    model.SKU = 1
		)

		s.cartRepository.DeleteItemMock.Expect(ctx, userID, sku).Return(nil)
		err := s.srv.DeleteItem(ctx, userID, sku)
		require.NoError(t, err)
	})
}
