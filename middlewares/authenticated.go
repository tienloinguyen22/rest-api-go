package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/rest-api-go/core/users"
	"github.com/tienloinguyen22/rest-api-go/utils"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get("user")
		if !exists {
			utils.HandleError(ctx, utils.NewApiError(http.StatusForbidden, "middlewares.authenticated.not-authenticated", errors.New("not authenticated")))
			return
		}

		_, ok := value.(*users.User)
		if !ok {
			utils.HandleError(ctx, utils.NewApiError(http.StatusForbidden, "middlewares.authenticated.not-authenticated", errors.New("not authenticated")))
			return
		}

		ctx.Next()
	}
}