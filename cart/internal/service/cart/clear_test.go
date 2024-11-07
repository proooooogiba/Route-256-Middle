package cart

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func (s *cartServiceTestSuite) TestClear() {
	s.T().Run("failed on repo.DeleteItemsByUserID", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		s.cartRepository.DeleteItemsByUserIDMock.Expect(ctx, userID).Return(ErrNoNil)
		err := s.srv.Clear(ctx, userID)
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
}
