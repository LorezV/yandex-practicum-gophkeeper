package v1

import (
	"github.com/gin-gonic/gin"

	usecase "github.com/LorezV/gophkeeper/internal/server/usecase"
	"github.com/LorezV/gophkeeper/pkg/logger"
)

type GophKeeperRoutes struct {
	uc usecase.GophKeeper
	l  logger.Interface
}

func newGophKeeperRoutes(handler *gin.RouterGroup, g usecase.GophKeeper, l logger.Interface) {
	r := &GophKeeperRoutes{g, l}

	handler.GET("/health", r.HealthCheck)

	userAPI := handler.Group("/user")
	{
		userAPI.GET("me", r.ProtectedByAccessToken(), r.UserInfo)

		userAPI.GET("logins", r.ProtectedByAccessToken(), r.GetLogins)
		userAPI.POST("logins", r.ProtectedByAccessToken(), r.AddLogin)
		userAPI.DELETE("logins/:id", r.ProtectedByAccessToken(), r.DelLogin)
		userAPI.PATCH("logins/:id", r.ProtectedByAccessToken(), r.UpdateLogin)

		userAPI.GET("cards", r.ProtectedByAccessToken(), r.GetCards)
		userAPI.POST("cards", r.ProtectedByAccessToken(), r.AddCard)
		userAPI.DELETE("cards/:id", r.ProtectedByAccessToken(), r.DelCard)
		userAPI.PATCH("cards/:id", r.ProtectedByAccessToken(), r.UpdateCard)

		userAPI.GET("notes", r.ProtectedByAccessToken(), r.GetNotes)
		userAPI.POST("notes", r.ProtectedByAccessToken(), r.AddNote)
		userAPI.DELETE("notes/:id", r.ProtectedByAccessToken(), r.DelNote)
		userAPI.PATCH("notes/:id", r.ProtectedByAccessToken(), r.UpdateNote)

		userAPI.GET("binary/:id", r.ProtectedByAccessToken(), r.DownloadBinary)
		userAPI.GET("binary", r.ProtectedByAccessToken(), r.GetBinaries)
		userAPI.POST("binary", r.ProtectedByAccessToken(), r.AddBinary)
		userAPI.DELETE("binary/:id", r.ProtectedByAccessToken(), r.DelBinary)
		userAPI.POST("binary/:id/meta", r.ProtectedByAccessToken(), r.AddBinaryMeta)
	}

	authAPI := handler.Group("/auth")
	{
		authAPI.POST("/register", r.SignUpUser)
		authAPI.POST("/login", r.SignInUser)
		authAPI.GET("/refresh", r.RefreshAccessToken)
		authAPI.GET("/logout", r.LogoutUser)
	}
}

// type historyResponse struct {
// 	History []entity.GophKeeper `json:"history"`
// }

// @Summary     Show history
// @Description Show all GophKeeper history
// @ID          history
// @Tags  	    GophKeeper
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /GophKeeper/history [get]
// func (r *GophKeeperRoutes) history(c *gin.Context) {
// 	GophKeepers, err := r.g.History(c.Request.Context())
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - history")
// 		errorResponse(c, http.StatusInternalServerError, "database problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, historyResponse{GophKeepers})
// }

// type doTranslateRequest struct {
// 	Source      string `json:"source"       binding:"required"  example:"auto"`
// 	Destination string `json:"destination"  binding:"required"  example:"en"`
// 	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
// }

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    GophKeeper
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up GophKeeper"
// @Success     200 {object} entity.GophKeeper
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /GophKeeper/do-translate [post]
// func (r *GophKeeperRoutes) doTranslate(c *gin.Context) {
// 	var request doTranslateRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	GophKeeper, err := r.g.Translate(
// 		c.Request.Context(),
// 		entity.GophKeeper{
// 			Source:      request.Source,
// 			Destination: request.Destination,
// 			Original:    request.Original,
// 		},
// 	)
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusInternalServerError, "GophKeeper service problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, GophKeeper)
// }
