package grpc

import (
	"context"

	"github.com/flametest/wallet-demo/internal/service"
	"github.com/flametest/wallet-demo/proto"
)

var _ proto.WalletDemoServiceServer = (*WalletServer)(nil)

type WalletServer struct {
	proto.UnimplementedWalletDemoServiceServer
	service.WalletService
}

func (w WalletServer) CreateWallet(ctx context.Context, req *proto.CreateWalletReq) (*proto.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletServer) GetWalletDetail(ctx context.Context, req *proto.GetWalletDetailReq) (*proto.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w WalletServer) WalletTransfer(ctx context.Context, req *proto.WalletTransferReq) (*proto.WalletTransferResp, error) {
	//TODO implement me
	panic("implement me")
}
