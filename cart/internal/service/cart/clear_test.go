package cart

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func (s *cartServiceTestSuite) TestClear() {
	s.T().Run("failed on repo.Clear", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		s.cartRepository.ClearMock.Expect(ctx, userID).Return(ErrNoNil)
		err := s.srv.Clear(ctx, userID)
		require.Error(t, err)
	})

	s.T().Run("successful", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		s.cartRepository.ClearMock.Expect(ctx, userID).Return(nil)
		err := s.srv.Clear(ctx, userID)
		require.NoError(t, err)
	})
}
