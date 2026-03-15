package service

import (
	"context"
	"sort"

	"github.com/flametest/vita/verrors"
	"github.com/flametest/vita/vgorm"
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
	TransferFund(ctx context.Context, req *dto.WalletTransferReq) error
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
	walletRepo := w.container.GetRepository().GetWalletRepo()
	wallet, err := walletRepo.GetByDisplayId(ctx, displayId)
	if err != nil && !verrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if verrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, verrors.NotFoundError("wallet not found")
	}
	return wallet, nil
}

func (w *walletServiceImpl) TransferFund(ctx context.Context, req *dto.WalletTransferReq) error {
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return verrors.BadRequestError("invalid amount format")
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return verrors.BadRequestError("amount must be greater than zero")
	}

	walletRepo := w.container.GetRepository().GetWalletRepo()
	fromWallet, err := walletRepo.GetByDisplayId(ctx, req.FromDisplayId)
	if err != nil {
		if verrors.Is(err, gorm.ErrRecordNotFound) {
			return verrors.NotFoundError("source wallet not found")
		}
		return err
	}

	// Check sufficient balance
	if fromWallet.Balance.LessThan(amount) {
		return verrors.BadRequestError("insufficient balance")
	}

	toWallet, err := walletRepo.GetByDisplayId(ctx, req.ToDisplayId)
	if err != nil {
		if verrors.Is(err, gorm.ErrRecordNotFound) {
			return verrors.NotFoundError("destination wallet not found")
		}
		return err
	}

	fromWallet.Balance = fromWallet.Balance.Sub(amount)
	toWallet.Balance = toWallet.Balance.Add(amount)

	// sort to prevent deadlocks
	wallets := []*model.Wallet{fromWallet, toWallet}
	sort.Slice(wallets, func(i, j int) bool {
		return wallets[i].Id < wallets[j].Id
	})

	err = walletRepo.DoInTx(func(tx vgorm.Tx) error {
		newRepo := w.container.GetRepository().GetWalletRepo(tx)
		for _, wallet := range wallets {
			if err := newRepo.UpdateWithVersion(ctx, wallet); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
