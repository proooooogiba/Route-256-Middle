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

func TestOrderCancel(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		orderRepoMock *mock.OrderRepoMock
		stockRepoMock *mock.StockRepoMock
	}
	type data struct {
		name    string
		id      int64
		prepare func(f *fields)
		wantErr error
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
			name: "failed on stockRepo.ReserveCancel",
			id:   1,
			prepare: func(f *fields) {
				f.orderRepoMock.GetOrderByIDMock.Expect(ctx, 1).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Status: string(model.AwaitingPayment),
					Items:  nil,
				}, nil)
				f.stockRepoMock.ReserveCancelMock.Expect(ctx, nil).Return(errorapp.ErrNoNil)
			},
			wantErr: errorapp.ErrNoNil,
		},
		{
			name: "failed on orderRepo.SetStatus",
			id:   1,
			prepare: func(f *fields) {
				f.orderRepoMock.GetOrderByIDMock.Expect(ctx, 1).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Status: string(model.AwaitingPayment),
					Items:  nil,
				}, nil)
				f.stockRepoMock.ReserveCancelMock.Expect(ctx, nil).Return(nil)
				f.orderRepoMock.SetStatusMock.Expect(ctx, 1, model.Cancelled).Return(errorapp.ErrNoNil)
			},
			wantErr: errorapp.ErrNoNil,
		},
		{
			name: "success",
			id:   1,
			prepare: func(f *fields) {
				f.orderRepoMock.GetOrderByIDMock.Expect(ctx, 1).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Status: string(model.AwaitingPayment),
					Items:  nil,
				}, nil)
				f.stockRepoMock.ReserveCancelMock.Expect(ctx, nil).Return(nil)
				f.orderRepoMock.SetStatusMock.Expect(ctx, 1, model.Cancelled).Return(nil)
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
			t.Parallel()
			tt.prepare(&fieldsForTableTest)
			err := service.OrderCancel(ctx, tt.id)
			if tt.wantErr != nil {
				require.Error(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
