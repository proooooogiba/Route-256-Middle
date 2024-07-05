package cart

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	errorapp "gitlab.ozon.dev/ipogiba/homework/cart/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/service/cart"
	"testing"
)

type RepoTestSuite struct {
	suite.Suite

	repo cart.Repository
}

func TestRepoTestSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}

func (s *RepoTestSuite) SetupTest() {
	s.repo = NewRepository()
}

func (s *RepoTestSuite) clean() {
	s.repo = NewRepository()
}

func (s *RepoTestSuite) TestAddItem() {
	s.T().Run("add 1 item", func(t *testing.T) {
		var (
			ctx             = context.Background()
			SKU   model.SKU = 1
			count uint16    = 1
			item            = model.Item{
				SKU:   SKU,
				Count: count,
			}
			userID int64 = 1
		)

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.Equal(t, errorapp.ErrNotFoundUser, err)
		assert.Nil(t, items)

		err = s.repo.AddItem(ctx, userID, item)
		assert.NoError(t, err)

		items, err = s.repo.GetItemsByUserID(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, item, items[0])
		s.clean()
	})

	s.T().Run("add same item", func(t *testing.T) {
		var (
			ctx             = context.Background()
			SKU   model.SKU = 1
			count uint16    = 1
			item            = model.Item{
				SKU:   SKU,
				Count: count,
			}
			userID int64 = 1
		)

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.Equal(t, errorapp.ErrNotFoundUser, err)
		assert.Nil(t, items)

		err = s.repo.AddItem(ctx, userID, item)
		assert.NoError(t, err)

		err = s.repo.AddItem(ctx, userID, item)
		assert.NoError(t, err)

		items, err = s.repo.GetItemsByUserID(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, item.SKU, items[0].SKU)
		assert.Equal(t, uint16(2), items[0].Count)
		s.clean()
	})

	s.T().Run("add diffrent item", func(t *testing.T) {
		var (
			ctx   = context.Background()
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

		items, err := s.repo.GetItemsByUserID(ctx, userID)
		assert.Equal(t, errorapp.ErrNotFoundUser, err)
		assert.Nil(t, items)

		err = s.repo.AddItem(ctx, userID, item1)
		assert.NoError(t, err)

		err = s.repo.AddItem(ctx, userID, item2)
		assert.NoError(t, err)

		items, err = s.repo.GetItemsByUserID(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(items))
		assert.Equal(t, item1, items[0])
		assert.Equal(t, item2, items[1])
		s.clean()
	})
}

func BenchmarkAddItem(b *testing.B) {
	ctx := context.Background()
	repo := &InMemoryRepository{
		items: make(map[int64]model.Cart),
	}
	userID := int64(1)
	item := model.Item{
		SKU:   12345,
		Count: 1,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := repo.AddItem(ctx, userID, item)
		if err != nil {
			b.Error(err)
		}
	}
}
