package profiles

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/middlewares"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type ProfileController struct {
	FirebaseAdmin *firebase.App
	UserRepo *users.UserRepository
	ProfileService *ProfileService
}

func NewProfileController(router *gin.Engine, firebaseAdmin *firebase.App, userRepo *users.UserRepository, profileService *ProfileService) {
	profileController := &ProfileController{
		FirebaseAdmin: firebaseAdmin,
		UserRepo: userRepo,
		ProfileService: profileService,
	}
	profileController.SetupRouter(router)
}

func (c ProfileController) SetupRouter(router *gin.Engine) {
	controller := router.Group("/api/profiles")

	controller.GET("/", middlewares.VerifyToken(c.FirebaseAdmin, c.UserRepo), func (ctx *gin.Context) {
		user, err := c.ProfileService.GetAuthenticatedUserProfile(ctx)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		ctx.JSON(200, user)
	})
}