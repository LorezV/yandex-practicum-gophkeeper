package v1

import (
	"errors"
	"net/http"

	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
)

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary     Register user
// @Description add new user
// @ID          register
// @Tags  	    Auth
// @Accept      json
// @Produce     json
// @Param       request body loginPayload true "Sing up new user"
// @Success     201 {object} entity.User
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /auth/register [post].
func (r *GophKeeperRoutes) SignUpUser(ctx *gin.Context) {
	var payload *loginPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	user, err := r.uc.SignUpUser(ctx, payload.Email, payload.Password)
	if err == nil {
		ctx.JSON(http.StatusCreated, user)

		return
	}

	if errors.Is(err, errs.ErrWrongEmail) || errors.Is(err, errs.ErrEmailAlreadyExists) {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

// @Summary     Login user
// @Description getting user JWT
// @ID          login
// @Tags  	    Auth
// @Accept      json
// @Produce     json
// @Param       request body loginPayload true "Sing in user"
// @Success     200 {object} entity.JWT
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /auth/login [post].
func (r *GophKeeperRoutes) SignInUser(ctx *gin.Context) {
	var payload *loginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	jwt, err := r.uc.SignInUser(ctx, payload.Email, payload.Password)

	if err == nil {
		ctx.SetCookie("access_token", jwt.AccessToken, jwt.AccessTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("refresh_token", jwt.RefreshToken, jwt.RefreshTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("logged_in", "true", jwt.AccessTokenMaxAge, "/", jwt.Domain, false, false)

		ctx.JSON(http.StatusOK, jwt)

		return
	}

	if errors.Is(err, errs.ErrWrongCredentials) {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

// @Summary     Refresh token
// @Description refresh access token
// @ID          refresh
// @Tags  	    Auth
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.JWT
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /auth/refresh [get].
func (r *GophKeeperRoutes) RefreshAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "refresh token has not been found")

		return
	}

	jwt, err := r.uc.RefreshAccessToken(ctx, refreshToken)

	if err == nil {
		ctx.SetCookie("access_token", jwt.AccessToken, jwt.AccessTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("logged_in", "true", jwt.AccessTokenMaxAge, "/", jwt.Domain, false, false)

		ctx.JSON(http.StatusOK, jwt)

		return
	}

	errorResponse(ctx, http.StatusBadRequest, err.Error())
}

// @Summary     Logout
// @Description dropping cookies
// @ID          logout
// @Tags  	    Auth
// @Success     200
// @Router      /auth/logout [get].
func (r *GophKeeperRoutes) LogoutUser(ctx *gin.Context) {
	domainName := r.uc.GetDomainName()
	ctx.SetCookie("access_token", "", -1, "/", domainName, false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", domainName, false, true)
	ctx.SetCookie("logged_in", "", -1, "/", domainName, false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
