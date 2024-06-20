package stock

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	errorapp "gitlab.ozon.dev/ipogiba/homework/loms/internal/errors"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/service/stock/mock"
	"testing"
)

func TestStockInfo(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		stockRepoMock *mock.StockRepoMock
	}
	type data struct {
		name              string
		sku               model.SKU
		prepare           func(f *fields)
		expectedAvailable uint64
		wantErr           error
	}

	testData := []data{
		{
			name: "failed on stockRepo.GetStockBySKU",
			sku:  1,
			prepare: func(f *fields) {
				f.stockRepoMock.GetStockBySKUMock.Expect(ctx, 1).Return(nil, errorapp.ErrNoNil)
			},
			expectedAvailable: 0,
			wantErr:           errorapp.ErrNoNil,
		},
		{
			name: "success",
			sku:  1,
			prepare: func(f *fields) {
				f.stockRepoMock.GetStockBySKUMock.Expect(ctx, 1).Return(&model.Stock{
					Sku:        1,
					TotalCount: 10,
					Reserved:   8,
				}, nil)
			},
			expectedAvailable: 2,
			wantErr:           nil,
		},
	}

	ctrl := minimock.NewController(t)

	fieldsForTableTest := fields{
		stockRepoMock: mock.NewStockRepoMock(ctrl),
	}

	service := NewStockService(fieldsForTableTest.stockRepoMock)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(&fieldsForTableTest)
			available, err := service.StockInfo(ctx, tt.sku)
			require.Equal(t, tt.expectedAvailable, available)
			if tt.wantErr != nil {
				require.Error(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
