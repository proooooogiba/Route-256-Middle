package loms

import (
	"context"

	"gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) OrderPay(ctx context.Context, req *loms.OrderPayRequest) (*loms.OrderPayResponse, error) {
	err := s.orderImpl.OrderPay(ctx, req.GetOrderId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &loms.OrderPayResponse{}, nil
}
