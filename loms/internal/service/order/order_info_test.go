package order

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	errorapp "route256/loms/internal/errors"
	"route256/loms/internal/model"
	"route256/loms/internal/service/order/mock"
	"testing"
)

func TestOrderInfo(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		orderRepoMock *mock.OrderRepoMock
		stockRepoMock *mock.StockRepoMock
	}
	type data struct {
		name          string
		id            int64
		prepare       func(f *fields)
		expectedOrder *model.Order
		wantErr       error
	}

	testData := []data{
		{
			name: "failed on orderRepo.GetOrderByID",
			id:   1,
			prepare: func(f *fields) {
				f.orderRepoMock.GetOrderByIDMock.Expect(ctx, 1).Return(nil, errorapp.ErrNoNil)
			},
			wantErr: errorapp.ErrNoNil,
		},
		{
			name: "failed on orderRepo.GetOrderByID - handle ErrOrderNotFound",
			id:   1,
			prepare: func(f *fields) {
				f.orderRepoMock.GetOrderByIDMock.Expect(ctx, 1).Return(nil, errorapp.ErrOrderNotFound)
			},
			wantErr: errorapp.ErrOrderNotFound,
		},
		{
			name: "success",
			id:   1,
			prepare: func(f *fields) {
				f.orderRepoMock.GetOrderByIDMock.Expect(ctx, 1).Return(
					&model.Order{
						ID: 1,
					}, nil)
			},
			expectedOrder: &model.Order{
				ID: 1,
			},
			wantErr: nil,
		},
	}

	ctrl := minimock.NewController(t)

	fieldsForTableTest := fields{
		orderRepoMock: mock.NewOrderRepoMock(ctrl),
		stockRepoMock: mock.NewStockRepoMock(ctrl),
	}

	service := NewOrderService(fieldsForTableTest.orderRepoMock, fieldsForTableTest.stockRepoMock)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(&fieldsForTableTest)
			order, err := service.OrderInfo(ctx, tt.id)
			require.Equal(t, tt.expectedOrder, order)
			if tt.wantErr != nil {
				require.Error(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
