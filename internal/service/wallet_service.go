package service

import (
	"context"

	"github.com/flametest/wallet-demo/internal/container"
	"github.com/flametest/wallet-demo/internal/infra/model"
	"github.com/flametest/wallet-demo/pkg/dto"
)

type WalletService interface {
	CreateWallet(ctx context.Context, req *dto.CreateWalletReq) (*model.Wallet, error)
	GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error)
	TransferFund(ctx context.Context, req *dto.WalletTransferRequest) error
}

type walletServiceImpl struct {
	container container.Container
}

func NewWalletService(container container.Container) WalletService {
	return &walletServiceImpl{container: container}
}

func (w walletServiceImpl) CreateWallet(ctx context.Context, req *dto.CreateWalletReq) (*model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w walletServiceImpl) GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w walletServiceImpl) TransferFund(ctx context.Context, req *dto.WalletTransferRequest) error {
	//TODO implement me
	panic("implement me")
}
