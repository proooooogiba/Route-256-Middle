package order

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/service/order/mock"
)

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		orderRepoMock *mock.OrderRepoMock
		stockRepoMock *mock.StockRepoMock
	}
	type data struct {
		name       string
		userID     int64
		items      []*model.Item
		prepare    func(f *fields)
		expectedID int64
		wantErr    error
	}

	testData := []data{
		{
			name:   "failed on stockRepo.Reserve",
			userID: 1,
			items:  nil,
			prepare: func(f *fields) {
				f.orderRepoMock.CreateOrderMock.Expect(ctx, 1, nil).Return(
					&model.Order{
						ID:     1,
						UserID: 1,
						Status: "new",
						Items:  nil,
					})
				f.stockRepoMock.ReserveMock.Expect(ctx, nil).Return(errorapp.ErrNoNil)
				f.orderRepoMock.SetStatusMock.Expect(ctx, 1, model.Failed).Return(nil)
			},
			expectedID: 0,
			wantErr:    errorapp.ErrNoNil,
		},
		{
			name:   "failed on orderRepo.SetStatus: failed status",
			userID: 1,
			items:  nil,
			prepare: func(f *fields) {
				f.orderRepoMock.CreateOrderMock.Expect(ctx, 1, nil).Return(
					&model.Order{
						ID:     1,
						UserID: 1,
						Status: "new",
						Items:  nil,
					})
				f.stockRepoMock.ReserveMock.Expect(ctx, nil).Return(errorapp.ErrNoNil)
				f.orderRepoMock.SetStatusMock.Expect(ctx, 1, model.Failed).Return(errorapp.ErrNoNil)
			},
			expectedID: 0,
			wantErr:    errorapp.ErrNoNil,
		},
		{
			name:   "failed on orderRepo.SetStatus: awaiting_payment status",
			userID: 1,
			items:  nil,
			prepare: func(f *fields) {
				f.orderRepoMock.CreateOrderMock.Expect(ctx, 1, nil).Return(
					&model.Order{
						ID:     1,
						UserID: 1,
						Status: "new",
						Items:  nil,
					})
				f.stockRepoMock.ReserveMock.Expect(ctx, nil).Return(nil)
				f.orderRepoMock.SetStatusMock.Expect(ctx, 1, model.AwaitingPayment).Return(errorapp.ErrNoNil)
			},
			expectedID: 0,
			wantErr:    errorapp.ErrNoNil,
		},
		{
			name:   "success order creating",
			userID: 1,
			items:  nil,
			prepare: func(f *fields) {
				f.orderRepoMock.CreateOrderMock.Expect(ctx, 1, nil).Return(
					&model.Order{
						ID:     1,
						UserID: 1,
						Status: "new",
						Items:  nil,
					})
				f.stockRepoMock.ReserveMock.Expect(ctx, nil).Return(nil)
				f.orderRepoMock.SetStatusMock.Expect(ctx, 1, model.AwaitingPayment).Return(nil)
			},
			expectedID: 1,
			wantErr:    nil,
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
			t.Parallel()
			tt.prepare(&fieldsForTableTest)
			id, err := service.CreateOrder(ctx, tt.userID, tt.items)
			require.Equal(t, tt.expectedID, id)

			if tt.wantErr != nil {
				require.Error(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
