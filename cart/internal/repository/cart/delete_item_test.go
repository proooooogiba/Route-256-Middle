package cart

import (
	"context"
	"github.com/stretchr/testify/assert"
	errorapp "route256/cart/internal/errors"
	"route256/cart/internal/model"
	"testing"
)

func (s *RepoTestSuite) TestDeleteItem() {
	s.T().Run("not found user", func(t *testing.T) {
		var (
			ctx              = context.Background()
			userID int64     = 1
			sku    model.SKU = 1
		)

		err := s.repo.DeleteItem(ctx, userID, sku)
		assert.Equal(t, errorapp.ErrNotFoundUser, err)
	})

	s.T().Run("add item and delete it", func(t *testing.T) {
		var (
			ctx              = context.Background()
			userID int64     = 1
			sku    model.SKU = 1
			item             = model.Item{
				SKU:   1,
				Count: 11,
			}
		)

		err := s.repo.AddItem(ctx, userID, item)
		assert.NoError(s.T(), err)

		err = s.repo.DeleteItem(ctx, userID, sku)
		assert.NoError(t, err)

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.NoError(t, err)
		assert.Empty(t, items)
	})

	s.T().Run("delete unexisted item", func(t *testing.T) {
		var (
			ctx                        = context.Background()
			userID           int64     = 1
			unexistedItemSKU model.SKU = 2
			item                       = model.Item{
				SKU:   1,
				Count: 11,
			}
		)

		err := s.repo.AddItem(ctx, userID, item)
		assert.NoError(s.T(), err)

		err = s.repo.DeleteItem(ctx, userID, unexistedItemSKU)
		assert.NoError(t, err)

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.NoError(t, err)

		assert.Len(t, items, 1)
		assert.Equal(t, item, items[0])
	})
}

func BenchmarkDeleteItem(b *testing.B) {
	ctx := context.Background()
	repo := &InMemoryRepository{
		items: make(map[int64]model.Cart),
	}
	userID := int64(1)
	sku := model.SKU(12345)

	repo.AddItem(ctx, userID, model.Item{SKU: sku, Count: 1})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := repo.DeleteItem(ctx, userID, sku)
		if err != nil {
			b.Error(err)
		}

		repo.AddItem(ctx, userID, model.Item{SKU: sku, Count: 1})
	}
}
