package auth

import (
	"database/sql"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type AuthService struct {
	FirebaseAdmin *firebase.App
	UserRepo *users.UserRepository
}

func NewAuthService(firebaseAdmin *firebase.App, userRepo *users.UserRepository) *AuthService {
	return &AuthService{
		FirebaseAdmin: firebaseAdmin,
		UserRepo: userRepo,
	}
}

func (s AuthService) SignIn(c *gin.Context, payload *SignInPayload) (*users.User, error) {
	authClient, err := s.FirebaseAdmin.Auth(c)
	if err != nil {
		return nil, utils.NewApiError(http.StatusInternalServerError, "auth.sign-in.cant-get-firebase-auth", err)
	}

	token, err := authClient.VerifyIDToken(c, payload.IdToken)
	if err != nil {
		return nil, utils.NewApiError(http.StatusInternalServerError, "auth.sign-in.cant-verify-id-token", err)
	}

	firebaseInfo, err := authClient.GetUser(c, token.UID)
	if err != nil {
		return nil, utils.NewApiError(http.StatusInternalServerError, "auth.sign-in.cant-get-firebase-user", err)
	}

	// Already exist user => Update firebaseId
	existedUser, err := s.UserRepo.FindByEmail(c, firebaseInfo.Email)
	if err != nil {
		return nil, utils.NewApiError(http.StatusInternalServerError, "auth.sign-in.cant-get-user-by-email", err)
	}
	if existedUser != nil {
		existedUser.FirebaseID = firebaseInfo.UID
		existedUser.SignupProvider = utils.GetSignupProvider(firebaseInfo)
		s.UserRepo.UpdateFirebaseInfoByID(c, existedUser.ID, existedUser)
		return existedUser, nil
	}

	// New user => Create user record in postgres
	newUser := users.User{
		ID: uuid.New(),
		FullName: firebaseInfo.DisplayName,
		Email: firebaseInfo.Email,
		PhoneNo: utils.NullString{
			NullString: sql.NullString{
				String: firebaseInfo.PhoneNumber,
			},
		},
		AvatarUrl: utils.NullString{
			NullString: sql.NullString{
				String: firebaseInfo.PhoneNumber,
			},
		},
		FirebaseID: firebaseInfo.UID,
		SignupProvider: utils.GetSignupProvider(firebaseInfo),
		BankTransferCode: utils.GetBankTransferCode(),
	}
	return s.UserRepo.Create(c, &newUser)
}