package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/adapters/inbound"
)

func main() {
	r := gin.Default()
	inbound.SetupRouter(r)
	r.Run() // listen and serve on port 8080
}