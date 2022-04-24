package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *AuthService
}

func NewAuthController(router *gin.Engine, authService *AuthService) {
	authController := &AuthController{
		AuthService: authService,
	}
	authController.SetupRouter(router)
}

func (c AuthController) SetupRouter(router *gin.Engine) {
	controller := router.Group("/api/auth")

	controller.POST("/sign-in", func (ctx *gin.Context) {
		var payload SignInPayload
		ctx.BindJSON(&payload)

		user, err := c.AuthService.SignIn(ctx, &payload)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, user)
	})
}