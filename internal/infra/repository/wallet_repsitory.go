package repository

import (
	"context"

	"github.com/flametest/vita/verrors"
	"github.com/flametest/vita/vgorm"
	"github.com/flametest/wallet-demo/internal/infra/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	vgorm.BaseRepo
	Create(ctx context.Context, wallet *model.Wallet) error
	GetByName(ctx context.Context, walletName string) (*model.Wallet, error)
	GetByDisplayId(ctx context.Context, displayId string) (*model.Wallet, error)
	UpdateWithVersion(ctx context.Context, wallet *model.Wallet) error
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

func (t *walletRepositoryImpl) UpdateWithVersion(ctx context.Context, wallet *model.Wallet) error {
	result := t.db.WithContext(ctx).
		Model(&model.Wallet{}).
		Where("id = ? AND version = ?", wallet.Id, wallet.Version).
		Updates(map[string]interface{}{
			"balance": wallet.Balance,
			"version": gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return verrors.ConflictError("optimistic lock failed: wallet has been modified by another transaction")
	}

	wallet.Version++

	return nil
}
