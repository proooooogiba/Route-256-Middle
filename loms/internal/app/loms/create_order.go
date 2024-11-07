package loms

import (
	"context"

	"gitlab.ozon.dev/ipogiba/homework/loms/internal/model"
	"gitlab.ozon.dev/ipogiba/homework/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateOrder(ctx context.Context, req *loms.CreateOrderRequest) (*loms.CreateOrderResponse, error) {
	items := make([]*model.Item, len(req.Items))
	for i, item := range req.GetItems() {
		items[i] = &model.Item{
			SKU:   model.SKU(item.GetSku()),
			Count: uint16(item.GetCount()),
		}
	}

	orderID, err := s.orderImpl.CreateOrder(ctx, req.UserId, items)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &loms.CreateOrderResponse{OrderId: orderID}, nil
}
