package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/adapters/inbound"
	"github.com/tienloinguyen22/edwork-api-go/adapters/outbound"
	"github.com/tienloinguyen22/edwork-api-go/configs"
)

func main() {
	// Env configs
	configs.InitializeConfigs()

	// Firebase admin
	outbound.InitializeFirebaseAdmin()

	// Router
	r := gin.Default()
	inbound.SetupRouter(r)

	// Start app
	r.Run(configs.Configs.ADDRESS)
}