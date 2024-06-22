package loms

import (
	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/ipogiba/homework/cart/pkg/api/loms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LomsService struct {
	client desc.LomsClient
}

func NewLomsServiceClient(grpcAddr string) (*LomsService, error) {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "create grpc connection")
	}

	client := desc.NewLomsClient(conn)

	return &LomsService{
		client: client,
	}, nil
}
