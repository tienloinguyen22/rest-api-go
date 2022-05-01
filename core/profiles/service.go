package profiles

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"

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

func (s ProfileService) GetUserProfile(ctx *gin.Context) (*users.User, error) {
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

func (s ProfileService) UpdateUserProfile(ctx *gin.Context, payload *UpdateUserProfilePayload) (*users.User, error) {
	value, exists := ctx.Get("user")
	if !exists {
		return nil, utils.NewApiError(http.StatusForbidden, "profiles.update-user-profile.not-authenticated", errors.New("not authenticated"))
	}

	authenticatedUser, ok := value.(*users.User)
	if !ok {
		return nil,	utils.NewApiError(http.StatusForbidden, "profiles.update-user-profile.not-authenticated", errors.New("not authenticated"))
	}

	user, err := s.UserRepo.FindByID(ctx, authenticatedUser.ID)
	if err != nil {
		return nil, utils.NewApiError(http.StatusInternalServerError, "profiles.update-user-profile.cant-get-user-by-id", err)
	}

	if payload.FullName != "" && payload.FullName != user.FullName {
		user.FullName = payload.FullName
	}
	if payload.Address != "" && payload.Address != user.Address.String {
		user.Address = utils.NullString{
			NullString: sql.NullString{
				String: payload.Address,
				Valid: true,
			},
		}
	}
	if payload.Grade > 0 && payload.Grade != user.Grade.Int64 {
		user.Grade = utils.NullInt64{
			NullInt64: sql.NullInt64{
				Int64: payload.Grade,
				Valid: true,
			},
		}
	}
	if payload.School != "" && payload.School != user.School.String {
		user.School = utils.NullString{
			NullString: sql.NullString{
				String: payload.School,
				Valid: true,
			},
		}
	}
	if utils.IsValidGender(payload.Gender) && payload.Gender != user.Gender.String {
		user.Gender = utils.NullString{
			NullString: sql.NullString{
				String: payload.Gender,
				Valid: true,
			},
		}
	}
	if utils.IsValidOwnerType(payload.OwnerType) && payload.OwnerType != user.OwnerType.String {
		user.OwnerType = utils.NullString{
			NullString: sql.NullString{
				String: payload.OwnerType,
				Valid: true,
			},
		}
	}
	if payload.PhoneNo != "" && payload.PhoneNo != user.PhoneNo.String {
		matched, err := regexp.MatchString(utils.REGEX_PHONE_NO, payload.PhoneNo)
		if err != nil {
			return nil, utils.NewApiError(http.StatusBadRequest, "profiles.update-user-profile.cant-validate-phone-no", err)
		}
		if !matched {
			return nil, utils.NewApiError(http.StatusBadRequest, "profiles.update-user-profile.invalid-phone-no", fmt.Errorf("invalid phone no"),)
		}
		user.PhoneNo = utils.NullString{
			NullString: sql.NullString{
				String: payload.PhoneNo,
				Valid: true,
			},
		}
	}
	if !payload.Dob.IsZero() && !payload.Dob.Equal(user.Dob.Time) {
		user.Dob = utils.NullTime{
			NullTime: sql.NullTime{
				Time: payload.Dob,
				Valid: true,
			},
		}
	}
	
	return s.UserRepo.UpdateUserInfoByID(ctx, user)
}