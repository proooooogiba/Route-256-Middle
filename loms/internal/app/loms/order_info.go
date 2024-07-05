package loms

import (
	"context"
	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) OrderInfo(ctx context.Context, req *loms.OrderInfoRequest) (*loms.OrderInfoResponse, error) {
	order, err := s.orderImpl.OrderInfo(ctx, req.GetOrderId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return repackOrderToProto(order), nil
}

func repackOrderToProto(in *model.Order) *loms.OrderInfoResponse {
	res := &loms.OrderInfoResponse{
		Status: in.Status,
		UserId: in.UserID,
	}

	items := make([]*loms.Item, len(in.Items))

	for i, item := range in.Items {
		items[i] = &loms.Item{
			Sku:   int32(item.SKU),
			Count: uint32(item.Count),
		}
	}

	res.Items = items

	return res
}
