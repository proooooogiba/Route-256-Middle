package cart

import (
	"context"
	"github.com/stretchr/testify/assert"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"testing"
)

func (s *RepoTestSuite) TestClear() {
	s.T().Run("not found user", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		err := s.repo.DeleteItemsByUserID(ctx, userID)
		assert.Equal(t, errorapp.ErrNotFoundUser, err)
	})

	s.T().Run("clear cart with items", func(t *testing.T) {
		var (
			ctx   = context.Background()
			err   error
			item1 = model.Item{
				SKU:   1,
				Count: 11,
			}
			item2 = model.Item{
				SKU:   2,
				Count: 22,
			}

			userID int64 = 1
		)

		err = s.repo.AddItem(ctx, userID, item1)
		assert.NoError(t, err)

		err = s.repo.AddItem(ctx, userID, item2)
		assert.NoError(t, err)

		err = s.repo.DeleteItemsByUserID(ctx, userID)
		assert.NoError(t, err)

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.Equal(t, errorapp.ErrNotFoundUser, err)
		assert.Nil(t, items)
	})
}
