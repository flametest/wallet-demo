package repository

import (
	"context"

	"github.com/flametest/wallet-demo/internal/infra/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *model.Wallet) error
	GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error)
}

type walletRepositoryImpl struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepositoryImpl{db: db}
}

func (t *walletRepositoryImpl) Create(ctx context.Context, wallet *model.Wallet) error {
	return t.db.WithContext(ctx).Create(wallet).Error
}

func (t *walletRepositoryImpl) GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error) {
	var wallet model.Wallet
	err := t.db.WithContext(ctx).Where("display_id = ?", displayId).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
