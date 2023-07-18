package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Root
// @Accept */*
// @Produce json
// @Success 200
// @Failure 500
// @Router /health [get].
func (r *GophKeeperRoutes) HealthCheck(ctx *gin.Context) {
	err := r.uc.HealthCheck()
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "connected"})
}
