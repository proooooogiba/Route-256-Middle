package cart

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"route256/cart/internal/model"
	"route256/cart/internal/service/cart/mock"
	"testing"
)

var ErrNoNil = errors.New("fail")

type cartServiceTestSuite struct {
	suite.Suite

	productService *mock.ProductServiceMock
	cartRepository *mock.RepositoryMock

	srv *Service
}

func TestCartServiceSuite(t *testing.T) {
	suite.Run(t, new(cartServiceTestSuite))
}

func (s *cartServiceTestSuite) SetupSuite() {
	ctrl := minimock.NewController(s.T())
	s.cartRepository = mock.NewRepositoryMock(ctrl)
	s.productService = mock.NewProductServiceMock(ctrl)
	s.srv = NewService(s.cartRepository, s.productService)
}

func (s *cartServiceTestSuite) TestAddItem() {
	s.T().Run("failed on productService.GetProductBySKU", func(t *testing.T) {
		ctx := context.Background()
		s.productService.GetProductBySKUMock.Expect(ctx, 1).Return(nil, ErrNoNil)

		err := s.srv.AddItem(ctx, 1, 1, 1)
		require.Error(t, err)
	})

	s.T().Run("failed on repo.AddItem", func(t *testing.T) {
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

		s.productService.GetProductBySKUMock.Expect(ctx, 1).Return(&model.Product{
			SKU:  1,
			Name: "product",
		}, nil)
		s.cartRepository.AddItemMock.Expect(ctx, userID, item).Return(ErrNoNil)

		err := s.srv.AddItem(ctx, userID, SKU, count)
		require.Error(t, err)
	})

	s.T().Run("successful", func(t *testing.T) {
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

		s.productService.GetProductBySKUMock.Expect(ctx, 1).Return(&model.Product{
			SKU:  1,
			Name: "product",
		}, nil)
		s.cartRepository.AddItemMock.Expect(ctx, userID, item).Return(nil)

		err := s.srv.AddItem(ctx, userID, SKU, count)
		require.NoError(t, err)
	})
}
