package middlewares

import (
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/rest-api-go/core/users"
	"github.com/tienloinguyen22/rest-api-go/utils"
)

type AuthorizationHeader struct {
	Bearer string `header:"Authorization"`
}

func VerifyToken(firebaseAdmin *firebase.App, userRepo *users.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := AuthorizationHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			utils.HandleError(ctx, err)
			return
		}

		if h.Bearer != "" {
			authClient, err := firebaseAdmin.Auth(ctx)
			if err != nil {
				utils.HandleError(ctx, utils.NewApiError(http.StatusInternalServerError, "middlewares.verify_token.cant-get-firebase-auth", err))
				return
			}

			jwt := strings.Replace(h.Bearer, "Bearer ", "", 1)
			token, err := authClient.VerifyIDToken(ctx, jwt)
			if err != nil {
				utils.HandleError(ctx, utils.NewApiError(http.StatusInternalServerError, "middlewares.verify_token.cant-verify-token", err))
				return
			}

			user, err := userRepo.FindByFirebaseID(ctx, token.UID)
			if err != nil {
				utils.HandleError(ctx, utils.NewApiError(http.StatusInternalServerError, "middlewares.verify_token.cant-find-user-by-id", err))
				return
			}

			ctx.Set("user", user)
		}

		ctx.Next()
	}
}