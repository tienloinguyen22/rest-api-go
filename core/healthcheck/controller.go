package healthcheck

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	controller := router.Group("/healthcheck")
	controller.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!!",
		})
	})
}

