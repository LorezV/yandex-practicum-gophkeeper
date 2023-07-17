package v1

import (
	"net/http"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary     Get user logins
// @Description fetching user logins
// @ID          get_logins
// @Tags  	    Logins
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Login
// @Success     204 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/logins [get].
func (r *GophKeeperRoutes) GetLogins(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userLogins, err := r.uc.GetLogins(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userLogins) == 0 {
		ctx.Status(http.StatusNoContent)

		return
	}

	ctx.JSON(http.StatusOK, userLogins)
}

// @Summary     Add user login
// @Description new user login
// @ID          add_login
// @Tags  	    Logins
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Param       request body entity.Login true "card for save"
// @Success     201 {object} entity.Login
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/logins [post].
func (r *GophKeeperRoutes) AddLogin(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	var payloadLogin *entity.Login

	if err := ctx.ShouldBindJSON(&payloadLogin); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	if err := r.uc.AddLogin(ctx, payloadLogin, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusAccepted, payloadLogin)
}

// @Summary     Delete user login
// @Description del user login
// @ID          del_login
// @Tags  	    Logins
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param       id   path string  true  "login ID"
// @Success     202
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/logins/{id} [delete].
func (r *GophKeeperRoutes) DelLogin(ctx *gin.Context) {
	loginUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	if err := r.uc.DelLogin(ctx, loginUUID, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}

// @Summary     Update user login
// @Description update user login
// @ID          update_login
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags  	    Logins
// @Accept      json
// @Produce     json
// @Param       id   path string  true  "Login ID"
// @Param       request body entity.Login true "card for update"
// @Success     202
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/logins/{id} [patch].
func (r *GophKeeperRoutes) UpdateLogin(ctx *gin.Context) {
	loginUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	var payloadLogin *entity.Login

	if err := ctx.ShouldBindJSON(&payloadLogin); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	payloadLogin.ID = loginUUID

	if err := r.uc.UpdateLogin(ctx, payloadLogin, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}
