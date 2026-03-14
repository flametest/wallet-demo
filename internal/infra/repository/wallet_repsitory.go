package repository

import (
	"context"

	"github.com/flametest/vita/vgorm"
	"github.com/flametest/wallet-demo/internal/infra/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	vgorm.BaseRepo
	Create(ctx context.Context, wallet *model.Wallet) error
	GetByName(ctx context.Context, walletName string) (*model.Wallet, error)
	GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error)
	Upsert(ctx context.Context, wallet *model.Wallet) error
}

type walletRepositoryImpl struct {
	db *gorm.DB
	vgorm.BaseRepo
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	baseRepo := vgorm.NewBaseRepo(db)
	return &walletRepositoryImpl{
		db:       db,
		BaseRepo: baseRepo,
	}
}

func (t *walletRepositoryImpl) Create(ctx context.Context, wallet *model.Wallet) error {
	return t.db.WithContext(ctx).Create(wallet).Error
}

func (t *walletRepositoryImpl) GetByName(ctx context.Context, walletName string) (*model.Wallet, error) {
	var wallet model.Wallet
	err := t.db.WithContext(ctx).Where("name = ?", walletName).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (t *walletRepositoryImpl) GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error) {
	var wallet model.Wallet
	err := t.db.WithContext(ctx).Where("display_id = ?", displayId).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (t *walletRepositoryImpl) Upsert(ctx context.Context, wallet *model.Wallet) error {
	return t.db.WithContext(ctx).Save(wallet).Error
}
