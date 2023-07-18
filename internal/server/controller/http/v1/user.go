package v1

import (
	"net/http"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
)

func (r *GophKeeperRoutes) getUserFromCtx(ctx *gin.Context) (user entity.User, err error) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		err = errs.ErrUnexpectedError

		return
	}

	return currentUser.(entity.User), nil
}

func (r *GophKeeperRoutes) UserInfo(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	ctx.JSON(http.StatusOK, currentUser)
}
