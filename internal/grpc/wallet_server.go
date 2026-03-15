package grpc

import (
	"context"

	"github.com/flametest/wallet-demo/internal/service"
	"github.com/flametest/wallet-demo/pkg/dto"
	"github.com/flametest/wallet-demo/proto"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.WalletDemoServiceServer = (*WalletServer)(nil)

type WalletServer struct {
	proto.UnimplementedWalletDemoServiceServer
	service.WalletService
}

func (w WalletServer) CreateWallet(ctx context.Context, req *proto.CreateWalletReq) (*proto.Wallet, error) {
	createWalletReq := &dto.CreateWalletReq{
		Name: req.Name,
	}
	validate := validator.New()
	if err := validate.Struct(createWalletReq); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	wallet, err := w.WalletService.CreateWallet(ctx, createWalletReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Wallet{
		Name:      wallet.Name,
		DisplayId: wallet.DisplayId,
		Balance:   wallet.Balance.String(),
	}, nil
}

func (w WalletServer) GetWalletDetail(ctx context.Context, req *proto.GetWalletDetailReq) (*proto.Wallet, error) {
	getWalletDetailReq := &dto.GetWalletDetailReq{
		DisplayID: req.DisplayId,
	}
	validate := validator.New()
	if err := validate.Struct(getWalletDetailReq); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	wallet, err := w.WalletService.GetByDisplayId(ctx, getWalletDetailReq.DisplayID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Wallet{
		Name:      wallet.Name,
		DisplayId: wallet.DisplayId,
		Balance:   wallet.Balance.String(),
	}, nil
}

func (w WalletServer) WalletTransfer(ctx context.Context, req *proto.WalletTransferReq) (*proto.WalletTransferResp, error) {
	walletTransferReq := &dto.WalletTransferReq{
		FromDisplayId: req.FromDisplayId,
		ToDisplayId:   req.ToDisplayId,
		Amount:        req.Amount,
	}
	validate := validator.New()
	if err := validate.Struct(walletTransferReq); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := w.WalletService.TransferFund(ctx, walletTransferReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.WalletTransferResp{
		Message: "Transfer Success",
	}, nil
}
