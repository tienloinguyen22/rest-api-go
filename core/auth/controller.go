package auth

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	controller := router.Group("/auth")

	controller.POST("/sign-in", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!!",
		})
	})
}