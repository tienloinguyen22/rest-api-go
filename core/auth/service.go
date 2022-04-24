package auth

import (
	"database/sql"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type SignInPayload struct {
	IdToken string `json:"idToken" binding:"required"`
}

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
		fmt.Printf("error getting Auth client: %v\n", err)
		return nil, err
	}

	token, err := authClient.VerifyIDToken(c, payload.IdToken)
	if err != nil {
		fmt.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}

	firebaseInfo, err := authClient.GetUser(c, token.UID)
	if err != nil {
		fmt.Printf("error getting firebase info: %v\n", err)
		return nil, err
	}

	// Already exist user => Update firebaseId
	existedUser, err := s.UserRepo.FindByEmail(c, firebaseInfo.Email)
	if err != nil {
		fmt.Println("Existed user error: ", err)
		return nil, err
	}
	if existedUser != nil {
		existedUser.FirebaseID = firebaseInfo.UID
		existedUser.SignupProvider = utils.GetSignupProvider(firebaseInfo)
		s.UserRepo.UpdateFirebaseInfoByID(c, existedUser.ID, existedUser)
		return existedUser, nil
	}

	// New user => Create user record in postgres
	newUser := users.User{
		FullName: firebaseInfo.DisplayName,
		Email: firebaseInfo.Email,
		PhoneNo: sql.NullString{
			String: firebaseInfo.PhoneNumber,
		},
		AvatarUrl: sql.NullString{
			String: firebaseInfo.PhoneNumber,
		},
		FirebaseID: firebaseInfo.UID,
		SignupProvider: utils.GetSignupProvider(firebaseInfo),
		BankTransferCode: utils.GetBankTransferCode(),
	}
	return s.UserRepo.Create(c, &newUser)
}