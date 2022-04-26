package auth

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type AuthController struct {
	FirebaseAdmin *firebase.App
	UserRepo *users.UserRepository
	AuthService *AuthService
}

func NewAuthController(router *gin.Engine, firebaseAdmin *firebase.App, userRepo *users.UserRepository, authService *AuthService) {
	authController := &AuthController{
		FirebaseAdmin: firebaseAdmin,
		UserRepo: userRepo,
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
			utils.HandleError(ctx, err)
		}

		ctx.JSON(200, user)
	})
}