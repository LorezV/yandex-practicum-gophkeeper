package v1

import (
	"errors"
	"net/http"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var errBinaryNameNotGiven = errors.New("binary name has not given")

// @Summary     Get user binary data
// @Description fetching user binary data
// @ID          get_binary
// @Tags  	    Binary
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Binary
// @Success     204
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/binary [get].
func (r *GophKeeperRoutes) GetBinaries(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userBinaries, err := r.uc.GetBinaries(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userBinaries) == 0 {
		ctx.Status(http.StatusNoContent)

		return
	}

	ctx.JSON(http.StatusOK, userBinaries)
}

// @Summary     Add user binary data
// @Description saving user binary data
// @ID          add_binary
// @Tags  	    Binary
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param file formData file true "Body with binary file"
// @Param       name    query     string  true  "name for entity"
// @Accept      json
// @Produce     json
// @Success     201 {object} []entity.Binary
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/binary [post].
func (r *GophKeeperRoutes) AddBinary(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}
	var binary entity.Binary
	if ctx.Query("name") == "" {
		errorResponse(ctx, http.StatusBadRequest, errBinaryNameNotGiven.Error())

		return
	}
	binary.Name = ctx.Query("name")

	file, err := ctx.FormFile("file")
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	binary.FileName = file.Filename
	if err = r.uc.AddBinary(ctx, &binary, file, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	ctx.JSON(http.StatusCreated, binary)
}

// @Summary     Download user binary data
// @Description downloading user binary data
// @ID          download_binary
// @Tags  	    Binary
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param       id   path string  true  "Binary ID"
// @Accept      json
// @Produce     json
// @Success     200 {file} binary
// @Success     204
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/binary/{id} [get].
func (r *GophKeeperRoutes) DownloadBinary(ctx *gin.Context) {
	binaryUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	filePath, err := r.uc.GetUserBinary(ctx, &currentUser, binaryUUID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusOK)

	ctx.File(filePath)
}

// @Summary     Delete user binary
// @Description del user file
// @ID          del_binary
// @Tags  	    Binary
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param       id   path string  true  "Binary ID"
// @Success     202
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/binary/{id} [delete].
func (r *GophKeeperRoutes) DelBinary(ctx *gin.Context) {
	binaryUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	if err := r.uc.DelUserBinary(ctx, &currentUser, binaryUUID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}

// @Summary     Add meta data for binary file
// @Description saving meta data
// @ID          add_binary_meta
// @Tags  	    Binary
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param       id   path string  true  "Binary ID"
// @Param       request body []entity.Meta true "meta for save"
// @Accept      json
// @Produce     json
// @Success     202 {object} []entity.Meta
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /user/binary/{id}/meta [post].
func (r *GophKeeperRoutes) AddBinaryMeta(ctx *gin.Context) {
	binaryUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}
	var payloadMeta []entity.Meta

	if err = ctx.ShouldBindJSON(&payloadMeta); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}
	binary, err := r.uc.AddBinaryMeta(ctx, &currentUser, binaryUUID, payloadMeta)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusCreated, binary.Meta)
}
