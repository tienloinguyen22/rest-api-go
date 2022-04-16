package auth

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/adapters/outbound"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type SignInPayload struct {
	IdToken string `json:"idToken" binding:"required"`
}

func getSignupProvider(provider interface{}) string {
	signupProvider := utils.SignupProvider.EMAIL
	if provider.ProviderID == "facebook.com" {
		signupProvider = utils.SignupProvider.FACEBOOK
	} else if provider.ProviderID == "google.com" {
		signupProvider = utils.SignupProvider.GOOGLE
	}

	return signupProvider
}

func SignIn(c *gin.Context, payload *SignInPayload) (*users.User, error) {
	auth, err := outbound.FirebaseApp.Auth(c)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return nil, err
	}

	token, err := auth.VerifyIDToken(c, payload.IdToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		return nil, err
	}

	firebaseInfo, err := auth.GetUser(c, token.UID)
	if err != nil {
		log.Fatalf("error getting firebase info: %v\n", err)
		return nil, err
	}

	signupProvider := getSignupProvider(firebaseInfo)
	fmt.Println(signupProvider)

	// Already exist user => Update firebaseId

	// New user => Create user record in postgres

	return nil, nil
}