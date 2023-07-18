package v1

import (
	"net/http"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary     Get user cards
// @Description fetching user cards
// @ID          get_cards
// @Tags  	    Cards
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Card
// @Success     204 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/cards [get].
func (r *GophKeeperRoutes) GetCards(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userCards, err := r.uc.GetCards(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userCards) == 0 {
		ctx.Status(http.StatusNoContent)

		return
	}

	ctx.JSON(http.StatusOK, userCards)
}

// @Summary     Add user card
// @Description new user card
// @ID          add_card
// @Tags  	    Cards
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Param       request body entity.Card true "card for save"
// @Success     201 {object} entity.Card
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/cards [post].
func (r *GophKeeperRoutes) AddCard(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	var payloadCard *entity.Card

	if err := ctx.ShouldBindJSON(&payloadCard); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	if err := r.uc.AddCard(ctx, payloadCard, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusAccepted, payloadCard)
}

// @Summary     Delete user card
// @Description del user card
// @ID          del_card
// @Tags  	    Cards
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param       id   path string  true  "Card ID"
// @Success     202
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/cards/{id} [delete].
func (r *GophKeeperRoutes) DelCard(ctx *gin.Context) {
	cardUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	if err := r.uc.DelCard(ctx, cardUUID, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}

// @Summary     Update user card
// @Description update user card
// @ID          update_card
// @Tags  	    Cards
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Param       id   path string  true  "Card ID"
// @Param       request body entity.Card true "card for update"
// @Success     202
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/cards/{id} [patch].
func (r *GophKeeperRoutes) UpdateCard(ctx *gin.Context) {
	cardUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	var payloadCard *entity.Card

	if err := ctx.ShouldBindJSON(&payloadCard); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	payloadCard.ID = cardUUID

	if err := r.uc.UpdateCard(ctx, payloadCard, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}
