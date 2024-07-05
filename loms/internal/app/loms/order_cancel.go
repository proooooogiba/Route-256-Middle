package loms

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) OrderCancel(ctx context.Context, req *loms.OrderCancelRequest) (*loms.OrderCancelResponse, error) {
	err := s.orderImpl.OrderCancel(ctx, req.GetOrderId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &loms.OrderCancelResponse{}, nil
}
