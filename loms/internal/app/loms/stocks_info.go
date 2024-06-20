package loms

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) StocksInfo(ctx context.Context, req *loms.StocksInfoRequest) (*loms.StocksInfoResponse, error) {
	stock, err := s.stocksImpl.StockInfo(ctx, model.SKU(req.GetSku()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &loms.StocksInfoResponse{
		Count: stock,
	}, nil
}
