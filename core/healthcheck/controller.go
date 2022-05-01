package healthcheck

import (
	"github.com/gin-gonic/gin"
)

type HealthcheckController struct {}

func NewHealthcheckController(router *gin.Engine) {
	authController := &HealthcheckController{}
	authController.SetupRouter(router)
}

func (c HealthcheckController) SetupRouter(router *gin.Engine) {
	controller := router.Group("/api/healthcheck")

	controller.GET("/", func (ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!!",
		})
		ctx.Abort()
	})
}

