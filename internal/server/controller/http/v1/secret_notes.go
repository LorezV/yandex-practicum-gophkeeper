package v1

import (
	"net/http"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary     Get user notes
// @Description fetching user notes
// @ID          notes_cards
// @Tags  	    Notes
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.SecretNote
// @Success     204 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/notes [get].
func (r *GophKeeperRoutes) GetNotes(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userNotes, err := r.uc.GetNotes(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userNotes) == 0 {
		ctx.Status(http.StatusNoContent)

		return
	}

	ctx.JSON(http.StatusOK, userNotes)
}

// @Summary     Add user note
// @Description new user note
// @ID          add_note
// @Tags  	    Notes
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Param       request body entity.SecretNote true "note for save"
// @Success     201 {object} entity.SecretNote
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/notes [post].
func (r *GophKeeperRoutes) AddNote(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	var payloadNote *entity.SecretNote

	if err := ctx.ShouldBindJSON(&payloadNote); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	if err := r.uc.AddNote(ctx, payloadNote, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusAccepted, payloadNote)
}

// @Summary     Delete user note
// @Description del user note
// @ID          del_note
// @Tags  	    Notes
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param       id   path string  true  "Note ID"
// @Success     201
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/notes/{id} [delete].
func (r *GophKeeperRoutes) DelNote(ctx *gin.Context) {
	noteUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	if err := r.uc.DelNote(ctx, noteUUID, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}

// @Summary     Update user note
// @Description update user note
// @ID          update_note
// @Tags  	    Notes
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Param       id   path string  true  "Note ID"
// @Param       request body entity.SecretNote true "card for update"
// @Success     202
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/notes/{id} [patch].
func (r *GophKeeperRoutes) UpdateNote(ctx *gin.Context) {
	noteUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	var payloadNote *entity.SecretNote

	if err := ctx.ShouldBindJSON(&payloadNote); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	payloadNote.ID = noteUUID

	if err := r.uc.UpdateNote(ctx, payloadNote, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}
