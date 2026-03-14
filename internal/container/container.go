package container

import (
	"github.com/flametest/vita/vgorm"
	"github.com/flametest/wallet-demo/internal/config"
	"github.com/flametest/wallet-demo/internal/infra/repository"
)

type Container interface {
	GetRepository() repository.Repository
}

type containerImpl struct {
	repository repository.Repository
}

func NewContainer(config *config.Config) (Container, error) {
	db, err := vgorm.NewDB(config.Datasource)
	if err != nil {
		return nil, err
	}
	repo := repository.NewRepository(db)
	return &containerImpl{repository: repo}, nil
}

func (c *containerImpl) GetRepository() repository.Repository {
	return c.repository
}
