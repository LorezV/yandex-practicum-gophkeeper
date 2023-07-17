// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// Swagger docs.
	_ "github.com/LorezV/gophkeeper/docs/server"

	usecase "github.com/LorezV/gophkeeper/internal/server/usecase"
	"github.com/LorezV/gophkeeper/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.GophKeeper) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/api/v1")
	{
		newGophKeeperRoutes(h, t, l)
	}
}
