package utils

import "firebase.google.com/go/auth"

func GetSignupProvider(provider *auth.UserRecord) string {
	signupProvider := SIGN_UP_PROVIDER_EMAIL
	if provider.ProviderID == "facebook.com" {
		signupProvider = SIGN_UP_PROVIDER_FACEBOOK
	} else if provider.ProviderID == "google.com" {
		signupProvider = SIGN_UP_PROVIDER_GOOGLE
	} else if provider.ProviderID == "apple.com" {
		signupProvider = SIGN_UP_PROVIDER_APPLE
	}

	return signupProvider
}