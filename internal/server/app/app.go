// Package app configures and runs application.
package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	config "github.com/LorezV/gophkeeper/config/server"
	v1 "github.com/LorezV/gophkeeper/internal/server/controller/http/v1"
	usecase "github.com/LorezV/gophkeeper/internal/server/usecase"
	"github.com/LorezV/gophkeeper/internal/server/usecase/repo"
	"github.com/LorezV/gophkeeper/pkg/cache"
	"github.com/LorezV/gophkeeper/pkg/httpserver"
	"github.com/LorezV/gophkeeper/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	gophKeeperRepo := repo.New(cfg.PG.URL, l)
	gophKeeperRepo.Migrate()

	defer gophKeeperRepo.ShutDown()
	// Use case
	GophKeeperUseCase := usecase.New(
		gophKeeperRepo,
		cfg,
		cache.New(cfg.Cache.DefaultExpiration, cfg.Cache.CleanupInterval),
		l,
	)

	var err error

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, GophKeeperUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
