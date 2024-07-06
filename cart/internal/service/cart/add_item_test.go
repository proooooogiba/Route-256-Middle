package cart

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/cart/internal/service/cart/mock"
	"go.uber.org/goleak"
	"testing"
)

var ErrNoNil = errors.New("fail")

type cartServiceTestSuite struct {
	suite.Suite

	productService *mock.ProductServiceMock
	cartRepository *mock.RepositoryMock
	lomsService    *mock.LomsServiceMock

	srv *Service
}

func TestCartServiceSuite(t *testing.T) {
	defer goleak.VerifyNone(t)
	suite.Run(t, new(cartServiceTestSuite))
}

func (s *cartServiceTestSuite) SetupSuite() {
	ctrl := minimock.NewController(s.T())
	s.cartRepository = mock.NewRepositoryMock(ctrl)
	s.productService = mock.NewProductServiceMock(ctrl)
	s.lomsService = mock.NewLomsServiceMock(ctrl)
	s.srv = NewService(s.cartRepository, s.productService, s.lomsService)
}

func (s *cartServiceTestSuite) TestAddItem() {
	s.T().Run("failed on productService.GetProductBySKU", func(t *testing.T) {
		ctx := context.Background()
		s.productService.GetProductBySKUMock.Expect(ctx, 1).Return(nil, ErrNoNil)

		err := s.srv.AddItem(ctx, 1, 1, 1)
		require.Error(t, err)
	})

	s.T().Run("failed on lomsService.StocksInfo", func(t *testing.T) {
		var (
			ctx              = context.Background()
			SKU    model.SKU = 1
			userID int64     = 1
		)

		s.productService.GetProductBySKUMock.Expect(ctx, 1).Return(&model.Product{
			SKU:  1,
			Name: "product",
		}, nil)
		s.lomsService.StocksInfoMock.Expect(ctx, SKU).Return(nil, ErrNoNil)

		err := s.srv.AddItem(ctx, userID, 1, 1)
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
			SKU:  SKU,
			Name: "product",
		}, nil)
		s.lomsService.StocksInfoMock.Expect(ctx, SKU).Return(&model.StockItem{
			Sku:   SKU,
			Count: 20,
		}, nil)
		s.cartRepository.AddItemMock.Expect(ctx, userID, item).Return(ErrNoNil)

		err := s.srv.AddItem(ctx, userID, SKU, count)
		require.Error(t, err)
	})

	s.T().Run("stocks are less then required", func(t *testing.T) {
		var (
			ctx              = context.Background()
			SKU    model.SKU = 1
			count  uint16    = 1
			userID int64     = 1
		)

		s.productService.GetProductBySKUMock.Expect(ctx, 1).Return(&model.Product{
			SKU:  SKU,
			Name: "product",
		}, nil)
		s.lomsService.StocksInfoMock.Expect(ctx, SKU).Return(&model.StockItem{
			Sku:   SKU,
			Count: 0,
		}, nil)

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
			SKU:  SKU,
			Name: "product",
		}, nil)
		s.lomsService.StocksInfoMock.Expect(ctx, SKU).Return(&model.StockItem{
			Sku:   SKU,
			Count: 10,
		}, nil)
		s.cartRepository.AddItemMock.Expect(ctx, userID, item).Return(nil)

		err := s.srv.AddItem(ctx, userID, SKU, count)
		require.NoError(t, err)
	})
}
