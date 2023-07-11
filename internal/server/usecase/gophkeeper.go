package usecase

import (
	config "github.com/LorezV/gophkeeper/config/server"
	c "github.com/LorezV/gophkeeper/pkg/cache"
	"github.com/LorezV/gophkeeper/pkg/logger"
)

const minutesPerHour = 60

// GophKeeperUseCase -.
type GophKeeperUseCase struct {
	repo  GophKeeperRepo
	cfg   *config.Config
	cache c.Cacher
	l     *logger.Logger
}

func New(r GophKeeperRepo, cfg *config.Config, cache c.Cacher, l *logger.Logger) *GophKeeperUseCase {
	return &GophKeeperUseCase{
		repo:  r,
		cfg:   cfg,
		cache: cache,
		l:     l,
	}
}

func (uc *GophKeeperUseCase) HealthCheck() error {
	return uc.repo.DBHealthCheck()
}

func (uc *GophKeeperUseCase) GetDomainName() string {
	return uc.cfg.Secutiry.Domain
}
