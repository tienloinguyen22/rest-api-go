package inbound

import (
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/core/healthcheck"
)

func SetupRouter(router *gin.Engine) {
	healthcheck.SetupRouter(router)
}