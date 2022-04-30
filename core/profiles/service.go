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
		return nil, utils.NewApiError(http.StatusForbidden, "profiles.get-user-profile.not-authenticated", errors.New("not authenticated"))
	}

	authenticatedUser, ok := value.(*users.User)
	if !ok {
		return nil,	utils.NewApiError(http.StatusForbidden, "profiles.get-user-profile.not-authenticated", errors.New("not authenticated"))
	}

	user, err := s.UserRepo.FindByID(ctx, authenticatedUser.ID)
	if err != nil {
		return nil, utils.NewApiError(http.StatusInternalServerError, "profiles.get-user-profile.cant-get-user-by-id", err)
	}

	return user, nil
}