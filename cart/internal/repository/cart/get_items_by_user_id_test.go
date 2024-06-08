package cart

import (
	"context"
	"github.com/stretchr/testify/assert"
	errorapp "route256/cart/internal/errors"
	"route256/cart/internal/model"
	"testing"
)

func (s *RepoTestSuite) TestGetItemsByUserID() {
	s.T().Run("not found user", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
		)

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.Equal(t, errorapp.ErrNotFoundUser, err)
		assert.Nil(t, items)
	})

	s.T().Run("add items to storage ", func(t *testing.T) {
		var (
			ctx          = context.Background()
			userID int64 = 1
			item         = model.Item{
				SKU:   1,
				Count: 11,
			}
		)

		err := s.repo.AddItem(ctx, userID, item)
		assert.NoError(t, err)

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, item, items[0])
	})
}
