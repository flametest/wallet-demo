package service

import (
	"context"

	"github.com/flametest/vita/verrors"
	"github.com/flametest/wallet-demo/internal/container"
	"github.com/flametest/wallet-demo/internal/infra/model"
	"github.com/flametest/wallet-demo/pkg/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
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

func (w *walletServiceImpl) CreateWallet(ctx context.Context, req *dto.CreateWalletReq) (*model.Wallet, error) {
	walletName := req.Name
	walletRepo := w.container.GetRepository().GetWalletRepo()
	wallet, err := walletRepo.GetByName(ctx, walletName)
	if err != nil && !verrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if wallet != nil {
		return nil, verrors.BadRequestError("wallet already exists")
	}
	newWallet := model.Wallet{
		Name:      walletName,
		DisplayId: uuid.New().String(),
		Balance:   decimal.Zero,
	}
	err = walletRepo.Create(ctx, &newWallet)
	if err != nil {
		return nil, err
	}
	return &newWallet, nil
}

func (w *walletServiceImpl) GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletServiceImpl) TransferFund(ctx context.Context, req *dto.WalletTransferRequest) error {
	//TODO implement me
	panic("implement me")
}
