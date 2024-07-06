package cart

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"sync"
	"testing"
)

func TestInMemoryRepository_Concurrency(t *testing.T) {
	t.Run("Concurrent AddItem", func(t *testing.T) {
		repo := NewRepository()
		userID := int64(1)
		itemCount := 100
		var wg sync.WaitGroup

		for i := 0; i < itemCount; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				err := repo.AddItem(context.Background(), userID, model.Item{SKU: model.SKU(i), Count: 1})
				assert.NoError(t, err)
			}(i)
		}

		wg.Wait()

		items, err := repo.GetItemsByUserID(context.Background(), userID)
		assert.NoError(t, err)
		assert.Len(t, items, itemCount)
	})

	t.Run("Concurrent AddItem and GetItemsByUserID", func(t *testing.T) {
		repo := NewRepository()
		userID := int64(1)
		itemCount := 100
		var wg sync.WaitGroup

		err := repo.AddItem(context.Background(), userID, model.Item{SKU: model.SKU(0), Count: 1})
		assert.NoError(t, err)

		// добавляем товары конкурентно
		for i := 0; i < itemCount; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				err := repo.AddItem(context.Background(), userID, model.Item{SKU: model.SKU(i), Count: 1})
				assert.NoError(t, err)
			}(i)
		}

		// пока добавляем товары, получаем товары пользователя конкуретно
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, err := repo.GetItemsByUserID(context.Background(), userID)
				assert.NoError(t, err)
			}()
		}

		wg.Wait()

		items, err := repo.GetItemsByUserID(context.Background(), userID)
		assert.NoError(t, err)
		assert.Len(t, items, itemCount)
	})

	t.Run("Concurrent AddItem for multiple users", func(t *testing.T) {
		repo := NewRepository()
		userCount := 10
		itemCount := 100
		var wg sync.WaitGroup

		for u := 0; u < userCount; u++ {
			for i := 0; i < itemCount; i++ {
				wg.Add(1)
				go func(u, i int) {
					defer wg.Done()
					err := repo.AddItem(context.Background(), int64(u), model.Item{SKU: model.SKU(i), Count: 1})
					assert.NoError(t, err)
				}(u, i)
			}
		}

		wg.Wait()

		for u := 0; u < userCount; u++ {
			items, err := repo.GetItemsByUserID(context.Background(), int64(u))
			assert.NoError(t, err)
			assert.Len(t, items, itemCount)
		}
	})
}
