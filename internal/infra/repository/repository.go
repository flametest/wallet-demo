package repository

import (
	"github.com/flametest/vita/vgorm"

	"gorm.io/gorm"
)

type Repository interface {
	GetWalletRepo(tx ...vgorm.Tx) WalletRepository
}

type repositoryImpl struct {
	db         *gorm.DB
	walletRepo WalletRepository
}

func (r repositoryImpl) GetWalletRepo(tx ...vgorm.Tx) WalletRepository {
	if len(tx) == 0 || tx[0] == nil {
		return r.walletRepo
	}
	t, ok := tx[0].(vgorm.Tx)
	if !ok {
		return r.walletRepo
	}
	return NewWalletRepository(t.DB())
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		walletRepo: NewWalletRepository(db),
	}
}
