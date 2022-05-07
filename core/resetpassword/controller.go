package resetpassword

import (
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/rest-api-go/utils"
	"gopkg.in/validator.v2"
)

type ResetPasswordController struct {
	ResetPasswordService *ResetPasswordService
}

func NewResetPasswordController(router *gin.Engine, resetPasswordService *ResetPasswordService) {
	resetPasswordController := &ResetPasswordController{
		ResetPasswordService: resetPasswordService,
	}
	resetPasswordController.SetupRouter(router)
}

func (c ResetPasswordController) SetupRouter(router *gin.Engine) {
	controller := router.Group("/api/reset-password")

	controller.POST("/request-token", func (ctx *gin.Context) {
		var payload RequestResetPasswordTokenPayload
		err := ctx.ShouldBindJSON(&payload)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		if err := validator.Validate(&payload); err != nil {
			utils.HandleError(ctx, err)
			return
		}

		err = c.ResetPasswordService.RequestResetPasswordToken(ctx, &payload)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		ctx.JSON(200, gin.H{
			"success": true,
		})
		ctx.Abort()
	})

	controller.GET("/verify-token/:token", func (ctx *gin.Context) {
		resetPasswordToken := ctx.Param("token")
		verified, err := c.ResetPasswordService.VerifyResetPasswordToken(ctx, resetPasswordToken)
		if err != nil {
			utils.HandleError(ctx, err)
			return
		}

		ctx.JSON(200, gin.H{
			"verified": verified,
		})
		ctx.Abort()
	})
}