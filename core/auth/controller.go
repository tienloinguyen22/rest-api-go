package auth

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/rest-api-go/core/users"
	"github.com/tienloinguyen22/rest-api-go/utils"
	"gopkg.in/validator.v2"
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
		err := ctx.ShouldBindJSON(&payload)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		if err := validator.Validate(&payload); err != nil {
			utils.HandleError(ctx, err)
			return
		}

		user, err := c.AuthService.SignIn(ctx, &payload)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		ctx.JSON(200, user)
		ctx.Abort()
	})
}