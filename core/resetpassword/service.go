package resetpassword

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/rest-api-go/adapters"
	"github.com/tienloinguyen22/rest-api-go/core/consumers"
	"github.com/tienloinguyen22/rest-api-go/core/users"
	"github.com/tienloinguyen22/rest-api-go/utils"
)

type ResetPasswordService struct {
	FirebaseAdmin *firebase.App
	MQ *adapters.MessageQueue
	ResetPasswordTokenRepo *ResetPasswordTokenRepository
	UserRepo *users.UserRepository
}

func NewResetPasswordService(
	firebaseAdmin *firebase.App,
	mq *adapters.MessageQueue,
	resetPasswordTokenRepo *ResetPasswordTokenRepository,
	userRepo *users.UserRepository,
) *ResetPasswordService {
	return &ResetPasswordService{
		FirebaseAdmin: firebaseAdmin,
		MQ: mq,
		ResetPasswordTokenRepo: resetPasswordTokenRepo,
		UserRepo: userRepo,
	}
}

func (s ResetPasswordService) RequestResetPasswordToken(ctx *gin.Context, payload *RequestResetPasswordTokenPayload) error {
	existedUser, err := s.UserRepo.FindByEmail(ctx, payload.Email)
	if err != nil {
		return utils.NewApiError(http.StatusInternalServerError, "resetpassword.request-reset-password-token.cant-get-user-by-email", err)
	}
	if existedUser == nil {
		return utils.NewApiError(http.StatusNotFound, "resetpassword.request-reset-password-token.user-not-found", err)
	}

	existedResetPasswordToken, err := s.ResetPasswordTokenRepo.FindNonExpiredByUserID(ctx, existedUser.ID)
	if err != nil {
		return utils.NewApiError(http.StatusInternalServerError, "resetpassword.request-reset-password-token.cant-get-existed-forgot-password-token", err)
	}
	if existedResetPasswordToken == nil {
		newResetPasswordToken := ResetPasswordToken{
			UserID: existedUser.ID,
			ExpiredAt: time.Now().Add(30 * time.Minute),
			CommonEntityFields: utils.CommonEntityFields{
				CreatedBy: existedUser.ID.String(),
			},
		}
		existedResetPasswordToken, err = s.ResetPasswordTokenRepo.Create(ctx, &newResetPasswordToken)
		if err != nil {
			return utils.NewApiError(http.StatusInternalServerError, "resetpassword.request-reset-password-token.cant-create-forgot-password-token", err)
		}
	}

	// Send email
	return s.MQ.Publish("SEND_MAIL", consumers.SendMailPayload{
		Template: "forgot-password.html",
		MailData: struct{
			Fullname string
			ResetPasswordLink string
		}{
			Fullname: existedUser.FullName,
			ResetPasswordLink: fmt.Sprintf("http://localhost:3000/reset-password/%v", existedResetPasswordToken.ID.String()),
		},
		To: []string{payload.Email},
	})
}

func (s ResetPasswordService) VerifyResetPasswordToken(ctx *gin.Context, resetPasswordToken string) (bool, error) {
	existedResetPasswordToken, err := s.ResetPasswordTokenRepo.FindNonExpiredByID(ctx, resetPasswordToken)
	if err != nil {
		return false, utils.NewApiError(http.StatusBadRequest, "resetpassword.verify-reset-password-token.cant-get-existed-forgot-password-token", err)
	}

	if existedResetPasswordToken == nil {
		return false, nil
	}

	return true, nil
}

func (s ResetPasswordService) ResetPassword(ctx *gin.Context, payload *ResetPasswordPayload) error {
	existedResetPasswordToken, err := s.ResetPasswordTokenRepo.FindNonExpiredByID(ctx, payload.ResetPasswordToken)
	if err != nil {
		return utils.NewApiError(http.StatusBadRequest, "resetpassword.reset-password.cant-get-existed-forgot-password-token", err)
	}
	if existedResetPasswordToken == nil {
		return utils.NewApiError(http.StatusBadRequest, "resetpassword.reset-password.cant-get-existed-forgot-password-token", errors.New("reset password token not found"))
	}

	existedUser, err := s.UserRepo.FindByID(ctx, existedResetPasswordToken.UserID)
	if err != nil {
		return utils.NewApiError(http.StatusBadRequest, "resetpassword.reset-password.cant-get-existed-user", err)
	}
	if existedUser == nil {
		return utils.NewApiError(http.StatusBadRequest, "resetpassword.reset-password.user-not-found", errors.New("user not found"))
	}

	authClient, err := s.FirebaseAdmin.Auth(ctx)
	if err != nil {
		return utils.NewApiError(http.StatusInternalServerError, "resetpassword.reset-password.cant-get-firebase-auth", err)
	}

	updateUserInfo := &auth.UserToUpdate{}
	updateUserInfo.Password(payload.NewPassword)
	_, err = authClient.UpdateUser(ctx, existedUser.FirebaseID, updateUserInfo)
	if err != nil {
		return utils.NewApiError(http.StatusInternalServerError, "resetpassword.reset-password.cant-update-user-password", err)
	}

	_, err = s.ResetPasswordTokenRepo.Expire(ctx, payload.ResetPasswordToken)
	if err != nil {
		utils.NewApiError(http.StatusInternalServerError, "resetpassword.reset-password.cant-expire-token", err)
	}

	return nil
}