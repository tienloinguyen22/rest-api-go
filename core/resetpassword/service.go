package resetpassword

import (
	"fmt"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/adapters"
	"github.com/tienloinguyen22/edwork-api-go/core/consumers"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/utils"
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