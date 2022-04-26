package profiles

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type ProfileService struct {
	UserRepo *users.UserRepository
}

func NewProfileService(userRepo *users.UserRepository) *ProfileService {
	return &ProfileService{
		UserRepo: userRepo,
	}
}

func (s ProfileService) GetAuthenticatedUserProfile(ctx *gin.Context) (*users.User, error) {
	value, exists := ctx.Get("user")
	if !exists {
		return nil, utils.NewApiError(http.StatusForbidden, "middlewares.authenticated.not-authenticated", errors.New("not authenticated"))
	}

	user, ok := value.(*users.User)
	if !ok {
		return nil,	utils.NewApiError(http.StatusForbidden, "middlewares.authenticated.not-authenticated", errors.New("not authenticated"))
	}

	return user, nil
}