package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, err error) {
	re, ok := err.(ApiError)
	if ok {
		fmt.Printf("%v: %v\n", re.Code, re.Err)
		ctx.JSON(re.Status, gin.H{
			"status": re.Status,
			"code": re.Code,
			"message": re.Err.Error(),
		})
	} else {
		fmt.Printf("internal server error: %v\n", re.Err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"code": "common.internal-server-error",
			"message": err.Error(),
		})
	}
}